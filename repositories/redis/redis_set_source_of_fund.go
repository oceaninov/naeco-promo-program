package redis

import (
	"context"
	"fmt"
	"github.com/oceaninov/naeco-promo-program/constants"
)

func (c *cache) RedisSetSourceOfFundBalance(ctx context.Context, accountNumber string, balance int64) error {
	key := fmt.Sprintf(constants.SourceOfFundsRedisKey, accountNumber)
	return c.Client.HMSet(ctx, key, map[string]interface{}{
		constants.BalanceRedisHashKey: balance,
	}).Err()
}
