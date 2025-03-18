package task

import "go-quickstart/api/middleware/taskchain"

type TestTask interface {
	InitContextData(args ...any) *taskchain.TaskContext
	TaskNode1(tc *taskchain.TaskContext)
	TaskNode2(tc *taskchain.TaskContext)
}
