package service

import (
	"github.com/covid-19-au/covid-19-api-server/database"
	u "github.com/covid-19-au/covid-19-api-server/util"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"golang.org/x/net/context"
	"time"
)

type AdminServer struct{}

type PublicServer struct{}

func (s AdminServer) IsServiceReady(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
	return isServiceReady()
}

func (s PublicServer) IsServiceReady(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
	return isServiceReady()
}

func isServiceReady() (*wrappers.BoolValue, error) {
	isMongoReady := database.PingMongo(5 * time.Second)
	isRedisReady := database.PingRedis()

	checkResult := isMongoReady && isRedisReady

	return u.WrapBool(checkResult), nil
}
