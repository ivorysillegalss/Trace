package consume

import (
	"github.com/thoas/go-funk"
	"go-quickstart/constant/mq"
	"go-quickstart/infrastructure/log"
	"go-quickstart/infrastructure/rabbitmq"
	"time"
)

type MessageHandler interface {
	PublishMessage(queueName string, data []byte)
	ConsumeMessage(queueName string, handler func([]byte) error)
	InitMessageQueue(args ...any)
}

var messageQueueConfigs map[string]MessageQueueArgs

func NewMessageHandler(conn *rabbitmq.Connection) MessageHandler {
	return &baseMessageHandler{conn: conn}
}

//TODO 进一步解耦MessageQueueArgs&InitMessageQueue与具体实现的关系

type baseMessageHandler struct {
	conn *rabbitmq.Connection
}

type MessageQueueArgs struct {
	ExchangeName    string
	QueueName       string
	KeyName         string
	ProducerChannel *rabbitmq.RabbitMqChannel
	ConsumerChannel *rabbitmq.RabbitMqChannel
}

func (b baseMessageHandler) InitMessageQueue(args ...any) {
	for _, argAny := range args {
		arg := argAny.(*MessageQueueArgs)

		// 创建生产者 Channel
		prodCh, err := b.conn.NewChannel()
		if err != nil {
			log.GetJsonLogger().WithFields("create producer channel err", err.Error()).Fatal(mq.PublishErr)
		} else {
			log.GetTextLogger().Info("create publish channel success")
		}
		// 创建消费者 Channel
		consCh, err := b.conn.NewChannel()
		if err != nil {
			log.GetJsonLogger().WithFields("create consumer channel err", err.Error()).Fatal(mq.ConsumeErr)
		} else {
			log.GetTextLogger().Info("create consumer channel success")
		}

		if err := prodCh.ExchangeDeclare(arg.ExchangeName, "direct"); err != nil {
			log.GetJsonLogger().WithFields("create exchange err", err.Error()).Fatal(mq.PublishErr)
		} else {
			log.GetTextLogger().Info("create exchange success")
		}

		if err := prodCh.QueueDeclare(arg.QueueName); err != nil {
			log.GetJsonLogger().WithFields("create queue err:", err.Error()).Fatal(mq.PublishErr)
		} else {
			log.GetTextLogger().Info("create queue success")
		}

		if err := prodCh.QueueBind(arg.QueueName, arg.KeyName, arg.ExchangeName); err != nil {
			log.GetJsonLogger().WithFields("bind queue err:", err.Error()).Fatal(mq.PublishErr)
		} else {
			log.GetTextLogger().Info("bind queue success")
		}

		// 保存创建的 Channel
		arg.ProducerChannel = prodCh
		arg.ConsumerChannel = consCh

		// 将配置保存到全局 map 中
		if funk.IsEmpty(messageQueueConfigs) {
			messageQueueConfigs = make(map[string]MessageQueueArgs)
		}
		messageQueueConfigs[arg.QueueName] = *arg
	}
}

func (b baseMessageHandler) PublishMessage(queueName string, data []byte) {
	args := messageQueueConfigs[queueName]
	go func() {
		if err := args.ProducerChannel.Publish(args.ExchangeName, args.KeyName, data); err != nil {
			log.GetJsonLogger().WithFields("publish msg err", err.Error()).Fatal(mq.PublishErr)
		}
		time.Sleep(time.Second)
	}()
}

func (b baseMessageHandler) ConsumeMessage(queueName string, handler func([]byte) error) {
	go func() {
		var ch *rabbitmq.RabbitMqChannel
		var err error

		// 尝试从 messageQueueConfigs 中获取现有的 Channel
		if config, exists := messageQueueConfigs[queueName]; exists && config.ConsumerChannel != nil {
			ch = config.ConsumerChannel
		} else {
			// 如果没有现有的 ConsumerChannel，则创建一个新的 Channel
			ch, err = b.conn.NewChannel()
			if err != nil {
				log.GetTextLogger().WithFields("new mq channel err", err.Error()).Fatal(mq.ConsumeErr)
				return
			}

			// 将新的 Channel 保存到 messageQueueConfigs 中
			messageQueueConfigs[queueName] = MessageQueueArgs{
				ExchangeName:    config.ExchangeName,
				QueueName:       queueName,
				KeyName:         config.KeyName,
				ConsumerChannel: ch,
				ProducerChannel: config.ProducerChannel, // 保持现有的 ProducerChannel 不变
			}
		}

		if err := b.conn.NewConsumer(ch, queueName, handler); err != nil {
			log.GetJsonLogger().WithFields("consume err", err.Error()).Fatal(mq.ConsumeErr)
		}
		log.GetTextLogger().Info("consume message success")
	}()
}
