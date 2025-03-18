package task

import (
	"go-quickstart/api/middleware/taskchain"
	"go-quickstart/bootstrap"
	"go-quickstart/constant/task"
	"go-quickstart/domain"
	"go-quickstart/infrastructure/log"
)

type testTask struct {
	testRepository domain.TestRepository
	env            *bootstrap.Env
	channels       *bootstrap.Channels
	poolFactory    *bootstrap.PoolsFactory
}

func (t *testTask) InitContextData(args ...any) *taskchain.TaskContext {
	log.GetTextLogger().Info("data init")
	return &taskchain.TaskContext{BusinessType: task.TestTaskType, BusinessCode: task.TestTaskCode}
}

func (t *testTask) TaskNode1(tc *taskchain.TaskContext) {
	log.GetTextLogger().Info("its node 1")
}

func (t *testTask) TaskNode2(tc *taskchain.TaskContext) {
	log.GetTextLogger().Info("its node 2")
	tc.TaskContextResponse = &taskchain.TaskContextResponse{
		Message: task.SuccessExecuteMessage,
		Code:    task.SuccessCode,
	}
}

func NewTestTask(ctx domain.TestRepository, env *bootstrap.Env, channels *bootstrap.Channels, factory *bootstrap.PoolsFactory) TestTask {
	return &testTask{testRepository: ctx, env: env, channels: channels, poolFactory: factory}
}
