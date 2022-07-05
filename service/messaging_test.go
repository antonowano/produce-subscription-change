package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type SuccessfulWriter struct {
	Writer
}

func (w *SuccessfulWriter) WriteMessages(_ context.Context, _ ...kafka.Message) error {
	return nil
}

type UnsuccessfulWriter struct {
	Writer
}

func (w *UnsuccessfulWriter) WriteMessages(_ context.Context, _ ...kafka.Message) error {
	return errors.New("some error")
}

type UnstableWriter struct {
	Writer
}

func (w *UnstableWriter) WriteMessages(_ context.Context, messages ...kafka.Message) error {
	writeErrors := make(kafka.WriteErrors, len(messages))

	for i := range writeErrors {
		if i%2 == 0 {
			writeErrors[i] = errors.New("some error")
		}
	}

	return writeErrors
}

func TestNewMessaging(t *testing.T) {
	writer := new(Writer)
	messaging := NewMessaging(*writer)
	assert.Equal(t, *writer, messaging.writer)
}

// Проверяем кейс, когда все сообщения успешно отправлены
func TestMessaging_SendMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	writer := &SuccessfulWriter{}
	messages := make([]Message, 0)
	messages = append(messages, Message{})
	messages = append(messages, Message{})
	messages = append(messages, Message{})
	countSentMessages := 0
	messaging := NewMessaging(writer)
	err := messaging.SendMessages(&messages, func(message *Message) {
		countSentMessages++
	})
	assert.Nil(t, err)
	assert.Equal(t, len(messages), countSentMessages)
}

// Проверяем кейс, когда ни одно сообщение не отправлено из-за ошибки
func TestMessaging_SendMessagesFail(t *testing.T) {
	writer := &UnsuccessfulWriter{}
	messages := make([]Message, 0)
	messages = append(messages, Message{})
	messages = append(messages, Message{})
	countSentMessages := 0
	messaging := NewMessaging(writer)
	err := messaging.SendMessages(&messages, func(message *Message) {
		countSentMessages++
	})
	assert.Error(t, err)
	assert.Equal(t, 0, countSentMessages)
}

// Проверяем кейс, когда произошла неполная отправка
func TestMessaging_SendMessagesUnstable(t *testing.T) {
	writer := &UnstableWriter{}
	messages := make([]Message, 0)
	messages = append(messages, Message{})
	messages = append(messages, Message{})
	messages = append(messages, Message{})
	messages = append(messages, Message{})
	messages = append(messages, Message{})
	countSentMessages := 0
	messaging := NewMessaging(writer)
	err := messaging.SendMessages(&messages, func(message *Message) {
		countSentMessages++
	})
	assert.Error(t, err)
	assert.True(t, countSentMessages > 0)
}
