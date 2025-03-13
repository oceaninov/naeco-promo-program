package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/scheduler"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
	"net/http"

	"gitlab.com/nbdgocean6/nobita-util/hdr"

	"gitlab.com/nbdgocean6/nobita-promo-program/gvars"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories/microservices"
	"gitlab.com/nbdgocean6/nobita-util/lgr"

	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/soheilhy/cmux"
	"gitlab.com/nbdgocean6/nobita-promo-program/constants"
	ep "gitlab.com/nbdgocean6/nobita-promo-program/endpoints"
	pb "gitlab.com/nbdgocean6/nobita-promo-program/protocs/api/v1"
	"gitlab.com/nbdgocean6/nobita-promo-program/repositories"
	"gitlab.com/nbdgocean6/nobita-promo-program/service"
	"gitlab.com/nbdgocean6/nobita-promo-program/transports"
	"gitlab.com/nbdgocean6/nobita-util/cb"
	"gitlab.com/nbdgocean6/nobita-util/dbc"
	"gitlab.com/nbdgocean6/nobita-util/vlt"

	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"github.com/joho/godotenv"
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

func main() {
	gvars.Log = lgr.Create(constants.ServiceName)

	level.Info(gvars.Log).Log(lgr.LogInfo, "service started")

	ctx := context.Background()
	defer ctx.Done()

	vaultUrl, vaultPath, vaultToken := GetVaultConfig()

	vault, err := vlt.NewVLT(vaultToken, vaultUrl, vaultPath)
	if err != nil {
		panic(err)
	}

	//reporter := httpreporter.NewReporter(vault.Get("/tracer_conf:url"))
	//localEndpoint, _ := zipkin.NewEndpoint(constants.ServiceName, "http://gcp-nb-sbox01:0")
	//exporter := oczipkin.NewExporter(reporter, localEndpoint)
	//trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	//trace.RegisterExporter(exporter)
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
			AuthHostAndPort:      "gcp-nb-sbox01:50010",
			WhitelistHostAndPort: "gcp-nb-sbox01:50011",
		}
		repoConf.SchedulerConf = scheduler.Config{
			SchedulerHostAndPort: "gcp-nb-sbox01:1929",
		}
	}

	CreateCredentialBasicAuthForSystem(
		vault.Get("/basic_auth_system:username"),
		vault.Get("/basic_auth_system:password"),
	)

	programRepo, err := repositories.NewRepositories(ctx, repoConf, trcr)
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

func GetVaultConfig() (vaultUrl string, path string, token string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	vaultUrl = os.Getenv("VAULT_URL")
	path = os.Getenv("VAULT_PATH")
	token = os.Getenv("VAULT_TOKEN")

	return
}
