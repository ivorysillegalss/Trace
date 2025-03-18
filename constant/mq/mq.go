package mq

// For Kafka
const (
	FirstOffset = "first"

	MaxGoroutine = 10

	TestConsumeId    = 1
	TestConsumeGroup = "testGroup"
	TestConsumeTopic = "testTopic"
	TestServiceName  = "testService"
)

// For Rabbit
const (
	// RabbitMqReconnectDelay reconnect after delay seconds
	RabbitMqReconnectDelay = 3

	PublishErr = "PUBLISH ERR"
	ConsumeErr = "CONSUME ERR"
)
