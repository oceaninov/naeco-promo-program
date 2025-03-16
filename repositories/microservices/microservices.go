package microservices

import (
	"context"
	authsvc "github.com/oceaninov/naeco-promo-auth/service/interface"
	"gitlab.com/nbdgocean6/nobita-promo-program/dial"
	whitelistsvc "github.com/oceaninov/naeco-promo-whitelist/service/interface"
	"google.golang.org/grpc"
)

type Microservices struct {
	AuthConnections []*grpc.ClientConn
	AuthSvc         authsvc.Service
	WhitelistSvc    whitelistsvc.Service
}

type Config struct {
	AuthHostAndPort      string
	WhitelistHostAndPort string
}

func NewMicroservice(ctx context.Context, conf Config) (*Microservices, error) {
	var connections []*grpc.ClientConn

	authSvc, authConn, err := dial.ConnectAuthService(ctx, conf.AuthHostAndPort)
	if err != nil {
		return nil, err
	}
	connections = append(connections, authConn)

	whitelistSvc, whitelistConn, err := dial.ConnectWhitelistService(ctx, conf.WhitelistHostAndPort)
	if err != nil {
		return nil, err
	}
	connections = append(connections, whitelistConn)

	return &Microservices{
		AuthConnections: connections,
		AuthSvc:         authSvc,
		WhitelistSvc:    whitelistSvc,
	}, nil
}
