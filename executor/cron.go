package executor

import "go-quickstart/domain"

type CronExecutor struct {
	testCron domain.TestCron
}

// SetupCron 启动定时任务
func (d *CronExecutor) SetupCron() {
	go d.testCron.TestCron()
}

func NewCronExecutor(cron domain.TestCron) *CronExecutor {
	return &CronExecutor{testCron: cron}
}
