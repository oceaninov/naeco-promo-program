package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oceaninov/naeco-promo-program/gvars"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BasicAuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.New("authentication required")
			}
			authData, ok := md["authorization"]
			if !ok {
				return nil, errors.New("authentication required")
			}
			var username, password string
			var basicAuth, data []string
			if basicAuth = strings.Split(authData[0], " "); len(basicAuth) != 2 {
				return nil, errors.New("authentication required")
			}
			if data = strings.Split(basicAuth[1], ":"); len(basicAuth) != 2 {
				return nil, errors.New("authentication required")
			}
			username = data[0]
			password = data[1]

			key := fmt.Sprintf("%s_%s", gvars.HashKeyMap, username)
			hashedPassword, ok := gvars.SyncMapHashStorage.Load(key)
			if !ok {
				return nil, errors.New("username not found")
			}
			err := bcrypt.CompareHashAndPassword([]byte(hashedPassword.(string)), []byte(password))
			if err != nil {
				return nil, err
			}
			return next(ctx, request)
		}
	}
}

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var resp interface{}
			req, _ := json.Marshal(request)
			defer func(begin time.Time) {
				level.Info(logger).Log(
					"err", err,
					"took", time.Since(begin),
					"request", string(req),
				)
			}(time.Now())
			resp, err = next(ctx, request)
			if err != nil {
				return nil, err
			}
			return resp, nil
		}
	}
}

func CircuitBreakerMiddleware(command string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var resp interface{}
			var logicErr error
			err = hystrix.Do(command, func() (err error) {
				resp, logicErr = next(ctx, request)
				return logicErr
			}, func(err error) error {
				return err
			})
			if logicErr != nil {
				return nil, logicErr
			}
			if err != nil {
				return nil, status.Error(
					codes.Unavailable,
					errors.New("service is busy or unavailable, please try again later").Error(),
				)
			}
			return resp, nil
		}
	}
}
