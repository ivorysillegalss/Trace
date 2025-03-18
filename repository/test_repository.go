package repository

import (
	"go-quickstart/domain"
	"go-quickstart/infrastructure/mysql"
	"go-quickstart/infrastructure/redis"
)

type testRepository struct {
	rcl   redis.Client
	mysql mysql.Client
}

func (t *testRepository) TestRepo() bool {
	//t.rcl.Set()
	//t.mysql.Gorm().Select()
	return true
}

func NewTestRepository(rcl redis.Client, m mysql.Client) domain.TestRepository {
	return &testRepository{rcl: rcl, mysql: m}
}
