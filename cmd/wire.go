//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"go-quickstart/bootstrap"
	"go-quickstart/consume"
	"go-quickstart/cron"
	"go-quickstart/executor"
	"go-quickstart/internal/tokenutil"
	"go-quickstart/repository"
	"go-quickstart/task"
	"go-quickstart/usecase"
)

var appSet = wire.NewSet(
	bootstrap.NewEnv,
	tokenutil.NewTokenUtil,
	bootstrap.NewDatabases,
	bootstrap.NewRedisDatabase,
	bootstrap.NewMysqlDatabase,
	bootstrap.NewMongoDatabase,
	bootstrap.NewPoolFactory,
	bootstrap.NewChannel,
	bootstrap.NewControllers,
	bootstrap.NewExecutors,
	bootstrap.NewKafkaConf,
	bootstrap.NewEsEngine,
	bootstrap.NewSearchEngine,
	bootstrap.NewRabbitConnection,

	consume.NewMessageHandler,
	consume.NewTestEvent,

	cron.NewTestCron,

	executor.NewCronExecutor,
	executor.NewConsumeExecutor,

	repository.NewTestRepository,
	usecase.NewTestUsecase,

	task.NewTestTask,

	wire.Struct(new(bootstrap.Application), "*"),
)

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	wire.Build(appSet)
	return &bootstrap.Application{}, nil
}
