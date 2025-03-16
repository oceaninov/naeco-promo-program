package scheduler

import (
	clockwerk "github.com/nightsilvertech/clockwerk/client"
	"github.com/oceaninov/naeco-promo-util/vlt"
	"strings"
)

type Config struct {
	SchedulerHostAndPort string
}

func NewSchedulerV2(config Config, vault vlt.VLT) (clockwerk.ClockwerkClient, error) {
	username := vault.Get("/scheduler_basic_auth:username")
	password := vault.Get("/scheduler_basic_auth:password")
	configs := strings.Split(config.SchedulerHostAndPort, ":")
	client, err := clockwerk.NewClockwerk(configs[0], configs[1], username, password)
	if err != nil {
		return nil, err
	}
	return client, nil
}
