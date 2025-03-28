package repository

import (
	"go-quickstart/domain"
	"go-quickstart/infrastructure/redis"
)

type messageRepository struct {
	rcl redis.Client
}

func (m *messageRepository) InsertMessage(reqMessage *domain.Message) bool {
	return false
}

func NewMessageRepository(rcl redis.Client) domain.MessageRepository {
	return &messageRepository{rcl: rcl}
}
