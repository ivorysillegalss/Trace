// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"go-quickstart/bootstrap"
	"go-quickstart/consume"
	"go-quickstart/cron"
	"go-quickstart/executor"
	"go-quickstart/repository"
	"go-quickstart/task"
	"go-quickstart/usecase"
)

// Injectors from wire.go:

// InitializeApp init application.
func InitializeApp() (*bootstrap.Application, error) {
	env := bootstrap.NewEnv()
	databases := bootstrap.NewDatabases(env)
	poolsFactory := bootstrap.NewPoolFactory()
	channels := bootstrap.NewChannel()
	controllers := bootstrap.NewControllers()
	client := bootstrap.NewRedisDatabase(env)
	mysqlClient := bootstrap.NewMysqlDatabase(env)
	testRepository := repository.NewTestRepository(client, mysqlClient)
	testCron := cron.NewTestCron(testRepository)
	cronExecutor := executor.NewCronExecutor(testCron)
	kafkaConf := bootstrap.NewKafkaConf(env)
	testConsume := consume.NewTestEvent(env, kafkaConf)
	consumeExecutor := executor.NewConsumeExecutor(testConsume)
	bootstrapExecutor := bootstrap.NewExecutors(cronExecutor, consumeExecutor)
	elasticsearchClient := bootstrap.NewEsEngine(env)
	searchEngine := bootstrap.NewSearchEngine(elasticsearchClient)
	application := &bootstrap.Application{
		Env:          env,
		Databases:    databases,
		PoolsFactory: poolsFactory,
		Channels:     channels,
		Controllers:  controllers,
		Executor:     bootstrapExecutor,
		SearchEngine: searchEngine,
		KafkaConf:    kafkaConf,
	}
	return application, nil
}

// wire.go:

var appSet = wire.NewSet(bootstrap.NewEnv, bootstrap.NewDatabases, bootstrap.NewRedisDatabase, bootstrap.NewMysqlDatabase, bootstrap.NewMongoDatabase, bootstrap.NewPoolFactory, bootstrap.NewChannel, bootstrap.NewControllers, bootstrap.NewExecutors, bootstrap.NewKafkaConf, bootstrap.NewEsEngine, bootstrap.NewSearchEngine, bootstrap.NewRabbitConnection, consume.NewMessageHandler, consume.NewTestEvent, cron.NewTestCron, executor.NewCronExecutor, executor.NewConsumeExecutor, repository.NewTestRepository, usecase.NewTestUsecase, task.NewTestTask, wire.Struct(new(bootstrap.Application), "*"))
