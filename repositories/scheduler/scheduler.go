package scheduler

import (
	"fmt"
	clockwerksvc "github.com/nightsilvertech/clockwerk/service/interface"
	clockwerktrpt "github.com/nightsilvertech/clockwerk/transports"
	grpcgoogle "google.golang.org/grpc"
)

type Config struct {
	SchedulerHostAndPort string
}

func NewScheduler(config Config) (clockwerksvc.Clockwerk, error) {
	dialOptions := []grpcgoogle.DialOption{
		grpcgoogle.WithInsecure(),
	}
	connectionString := fmt.Sprintf(config.SchedulerHostAndPort)
	conn, err := grpcgoogle.Dial(connectionString, dialOptions...)
	if err != nil {
		return nil, err
	}
	clockwerk := clockwerktrpt.ClockwerkClient(conn)
	return clockwerk, nil
}
