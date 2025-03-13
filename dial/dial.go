package dial

import (
	"context"
	authsvc "gitlab.com/nbdgocean6/nobita-promo-auth/service/interface"
	clientauth "gitlab.com/nbdgocean6/nobita-promo-auth/transports"
	whitelistsvc "gitlab.com/nbdgocean6/nobita-promo-whitelist/service/interface"
	clientwhitelist "gitlab.com/nbdgocean6/nobita-promo-whitelist/transports"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	grpcgoogle "google.golang.org/grpc"
)

func ConnectWhitelistService(ctx context.Context, hostAndPort string) (whitelistsvc.Service, *grpcgoogle.ClientConn, error) {
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		return nil, nil, err
	}
	dialOptions := []grpcgoogle.DialOption{
		grpcgoogle.WithInsecure(),
		grpcgoogle.WithStatsHandler(new(ocgrpc.ClientHandler)),
	}
	conn, err := grpcgoogle.DialContext(ctx, hostAndPort, dialOptions...)
	if err != nil {
		panic(err)
	}
	return clientwhitelist.NewGRPWhitelistClient(conn), conn, nil
}

func ConnectAuthService(ctx context.Context, hostAndPort string) (authsvc.Service, *grpcgoogle.ClientConn, error) {
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		return nil, nil, err
	}
	dialOptions := []grpcgoogle.DialOption{
		grpcgoogle.WithInsecure(),
		grpcgoogle.WithStatsHandler(new(ocgrpc.ClientHandler)),
	}
	conn, err := grpcgoogle.DialContext(ctx, hostAndPort, dialOptions...)
	if err != nil {
		panic(err)
	}
	return clientauth.NewGRPAuthClient(conn), conn, nil
}
