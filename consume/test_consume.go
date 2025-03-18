package consume

import (
	"context"
	"fmt"
	"go-quickstart/bootstrap"
	"go-quickstart/constant/mq"
	"go-quickstart/domain"
	kq "go-quickstart/infrastructure/kafka"
	"go-quickstart/infrastructure/log"
)

type testEvent struct {
	env            *bootstrap.Env
	kafka          *bootstrap.KafkaConf
	testRepository domain.TestRepository
}

type testHandler struct {
	testRepository domain.TestRepository
}

func (t *testHandler) Consume(ctx context.Context, key, value string) error {
	// with specific logic
	_ = t.testRepository.TestRepo()
	return nil
}

func (t *testEvent) TestConsumeWithFunc() {
	th := &testHandler{testRepository: t.testRepository}

	conf := *(t.kafka)
	kqConf := conf.ConfMap[mq.TestConsumeId]

	queue, err := kq.NewQueue(kqConf, th)
	if err != nil {
		log.GetTextLogger().Error(fmt.Sprintf("queue get error :" + err.Error()))
	}
	go queue.Start()
}

func (t *testEvent) TestConsume() {

	conf := *(t.kafka)
	kqConf := conf.ConfMap[mq.TestConsumeId]

	handler := kq.WithHandle(func(ctx context.Context, key, value string) error {
		_ = t.testRepository.TestRepo()
		// with specific logic
		return nil
	})

	queue, err := kq.NewQueue(kqConf, handler)
	if err != nil {
		log.GetTextLogger().Error(fmt.Sprintf("queue get error :" + err.Error()))
	}
	go queue.Start()
}

func NewTestEvent(env *bootstrap.Env, conf *bootstrap.KafkaConf) domain.TestConsume {
	return &testEvent{env: env, kafka: conf}
}
