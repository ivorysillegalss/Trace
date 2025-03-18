package cron

import (
	"go-quickstart/domain"
	"go-quickstart/infrastructure/log"
)

type testCron struct {
	testRepository domain.TestRepository
}

func (t *testCron) TestCron() {
	for {
		log.GetTextLogger().Info("cron starting")
	}
}

func NewTestCron(tr domain.TestRepository) domain.TestCron {
	return &testCron{testRepository: tr}
}
