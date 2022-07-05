package service

import (
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// Вызывает NewWriter и проверяет, что имя читаемого топика subscription
func TestNewWriter(t *testing.T) {
	addresses := strings.Split("A,B,C", ",")
	username := "username"
	password := "password"
	writer := NewWriter(addresses, username, password)
	kafkaWriter := writer.(*kafka.Writer)
	assert.Equal(t, "subscription", kafkaWriter.Topic, "Топик указан неправильно")
}

func TestWriter_Logf(t *testing.T) {
	writer := NewWriter([]string{}, "username", "password")
	kafkaWriter := writer.(*kafka.Writer)
	kafkaWriter.Logger.Printf("")
}
