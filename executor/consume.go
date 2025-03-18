package executor

import "go-quickstart/domain"

type ConsumeExecutor struct {
	testConsume domain.TestConsume
}

func (d *ConsumeExecutor) SetupConsume() {
	d.testConsume.TestConsume()
	d.testConsume.TestConsumeWithFunc()
}

func NewConsumeExecutor(tc domain.TestConsume) *ConsumeExecutor {
	return &ConsumeExecutor{testConsume: tc}
}
