package bootstrap

import (
	"go-quickstart/api/controller"
	"go-quickstart/executor"
	"go-quickstart/infrastructure/elasticsearch"
	"go-quickstart/infrastructure/mongo"
	"go-quickstart/infrastructure/mysql"
	"go-quickstart/infrastructure/pool"
	"go-quickstart/infrastructure/redis"
)

type Application struct {
	Env          *Env
	Databases    *Databases
	PoolsFactory *PoolsFactory
	Channels     *Channels
	Controllers  *Controllers
	Executor     *Executor
	SearchEngine *SearchEngine
	KafkaConf    *KafkaConf
}

type Databases struct {
	Mongo mongo.Client
	Redis redis.Client
	Mysql mysql.Client
}

// PoolsFactory k为pool业务号 v为poll详细配置信息
type PoolsFactory struct {
	Pools map[int]*pool.Pool
}

type Controllers struct {
	TestController controller.TestController
}

type Executor struct {
	CronExecutor    *executor.CronExecutor
	ConsumeExecutor *executor.ConsumeExecutor
}

type SearchEngine struct {
	EsEngine elasticsearch.Client
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Databases.Mongo)
}

type Channels struct {
	Stop chan bool
}

func NewControllers() *Controllers {
	return &Controllers{}
}

func NewExecutors(ce *executor.CronExecutor, cse *executor.ConsumeExecutor) *Executor {
	return &Executor{
		CronExecutor:    ce,
		ConsumeExecutor: cse,
	}
}
