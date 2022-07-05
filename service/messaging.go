package service

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
)

//go:generate mockgen -source=messaging.go -destination=../mocks/service/messaging.go

type Messaging interface {
	SendMessages(messages *[]Message, success func(*Message)) error
}

func NewMessaging(writer Writer) *MessagingImpl {
	return &MessagingImpl{
		writer: writer,
	}
}

type MessagingImpl struct {
	writer Writer
}

func (service *MessagingImpl) SendMessages(messages *[]Message, success func(*Message)) error {
	kafkaMessages := make([]kafka.Message, 0)

	for i := range *messages {
		jsonData, _ := (*messages)[i].ToJson()
		kafkaMessages = append(kafkaMessages, kafka.Message{
			Key:   []byte("subscription-" + (*messages)[i].ChangeId),
			Value: jsonData,
		})
	}

	switch err := service.writer.WriteMessages(context.Background(), kafkaMessages...).(type) {
	case nil:
		for i := range *messages {
			success(&(*messages)[i])
		}
		return nil
	case kafka.WriteErrors:
		for i := range *messages {
			if err[i] == nil {
				success(&(*messages)[i])
			}
		}
		return errors.New("отправка одного или нескольких сообщений вернула ошибку")
	default:
		return err
	}
}
