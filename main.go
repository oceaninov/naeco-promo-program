package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/oceaninov/naeco-promo-program/repositories/scheduler"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
	"net/http"

	"github.com/oceaninov/naeco-promo-util/hdr"

	"github.com/oceaninov/naeco-promo-program/gvars"
	"github.com/oceaninov/naeco-promo-program/repositories/microservices"
	"github.com/oceaninov/naeco-promo-util/lgr"

	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/soheilhy/cmux"
	"github.com/oceaninov/naeco-promo-program/constants"
	ep "github.com/oceaninov/naeco-promo-program/endpoints"
	pb "github.com/oceaninov/naeco-promo-program/protocs/api/v1"
	"github.com/oceaninov/naeco-promo-program/repositories"
	"github.com/oceaninov/naeco-promo-program/service"
	"github.com/oceaninov/naeco-promo-program/transports"
	"github.com/oceaninov/naeco-promo-util/cb"
	"github.com/oceaninov/naeco-promo-util/dbc"
	"github.com/oceaninov/naeco-promo-util/vlt"

	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// ServeGRPC serving GRPC server and will be merged using MergeServer function
func ServeGRPC(listener net.Listener, service pb.ProgramServiceServer, serverOptions []grpc.ServerOption) error {
	level.Info(gvars.Log).Log(lgr.LogInfo, "initialize grpc server")

	var grpcServer *grpc.Server
	if len(serverOptions) > 0 {
		grpcServer = grpc.NewServer(serverOptions...)
	} else {
		grpcServer = grpc.NewServer()
	}
	pb.RegisterProgramServiceServer(grpcServer, service)
	return grpcServer.Serve(listener)
}

// ServeHTTP serving HTTP server and will be merged using MergeServer function
func ServeHTTP(listener net.Listener, service pb.ProgramServiceServer) error {
	level.Info(gvars.Log).Log(lgr.LogInfo, "initialize rest server")

	mux := runtime.NewServeMux()
	err := pb.RegisterProgramServiceHandlerServer(context.Background(), mux, service)
	if err != nil {
		return err
	}
	srv := &http.Server{Handler: hdr.CORS(mux)}
	return srv.Serve(listener)
}

// MergeServer start ServeGRPC and ServeHTTP concurrently
func MergeServer(service pb.ProgramServiceServer, serverOptions []grpc.ServerOption) {
	port := fmt.Sprintf(":%s", "50012")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings(
		"content-type", "application/grpc",
	))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return ServeGRPC(grpcListener, service, serverOptions) })
	g.Go(func() error { return ServeHTTP(httpListener, service) })
	g.Go(func() error { return m.Serve() })

	log.Fatal(g.Wait())
}

func CreateCredentialBasicAuthForSystem(username, password string) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	key := fmt.Sprintf("%s_%s", gvars.HashKeyMap, username)
	gvars.SyncMapHashStorage.Store(key, string(hashByte))
}

func ServerAddress(host, port string) string {
	return fmt.Sprintf("http://%s:%s", host, port)
}

func ServerGrpcAddress(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func main() {
	const serverHost = `gcp-nb-sbox01`

	gvars.Log = lgr.Create(constants.ServiceName)

	level.Info(gvars.Log).Log(lgr.LogInfo, "service started")

	ctx := context.Background()
	defer ctx.Done()

	vault, err := vlt.NewVLT("myroot", ServerAddress(serverHost, "8300"), "/secret")
	if err != nil {
		panic(err)
	}

	reporter := httpreporter.NewReporter(vault.Get("/tracer_conf:url"))
	localEndpoint, _ := zipkin.NewEndpoint(constants.ServiceName, ServerAddress(serverHost, "0"))
	exporter := oczipkin.NewExporter(reporter, localEndpoint)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
	trcr := trace.DefaultTracer

	var repoConf repositories.RepoConf
	{
		repoConf.SQL = dbc.Config{
			Username: vault.Get("/sql_database:username"),
			Password: vault.Get("/sql_database:password"),
			Host:     vault.Get("/sql_database:host"),
			Port:     vault.Get("/sql_database:port"),
			Name:     vault.Get("/sql_database:db"),
		}
		repoConf.MicroserviceConf = microservices.Config{
			AuthHostAndPort:      ServerGrpcAddress(serverHost, "50010"),
			WhitelistHostAndPort: ServerGrpcAddress(serverHost, "50011"),
		}
		repoConf.SchedulerConf = scheduler.Config{
			SchedulerHostAndPort: ServerGrpcAddress(serverHost, "1929"),
		}
	}

	CreateCredentialBasicAuthForSystem(
		vault.Get("/basic_auth_system:username"),
		vault.Get("/basic_auth_system:password"),
	)

	programRepo, err := repositories.NewRepositories(ctx, repoConf, trcr, vault)
	if err != nil {
		panic(err)
	}

	err = cb.StartHystrix(10, constants.ServiceName)
	if err != nil {
		panic(err)
	}

	whitelistSvc := service.NewUsecases(*programRepo, trcr, vault)

	authEp, err := ep.NewProgramEndpoint(whitelistSvc, programRepo.Microservice.AuthSvc, gvars.Log, vault)
	if err != nil {
		panic(err)
	}

	server := transports.NewProgramServer(authEp)
	MergeServer(server, nil)
}
