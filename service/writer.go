package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

//go:generate mockgen -source=writer.go -destination=../mocks/service/writer.go

type Writer interface {
	WriteMessages(ctx context.Context, messages ...kafka.Message) error
	Close() error
}

func NewWriter(addresses []string, username string, password string) Writer {
	mechanism := plain.Mechanism{
		Username: username,
		Password: password,
	}
	transport := &kafka.Transport{
		SASL: mechanism,
	}
	logf := func(msg string, a ...interface{}) {
		fmt.Printf(msg, a...)
		fmt.Println()
	}
	return &kafka.Writer{
		Addr:         kafka.TCP(addresses...),
		Topic:        "subscription",
		Transport:    transport,
		Balancer:     &kafka.Hash{},
		RequiredAcks: kafka.RequireOne,
		Logger:       kafka.LoggerFunc(logf),
		ErrorLogger:  kafka.LoggerFunc(logf),
	}
}
