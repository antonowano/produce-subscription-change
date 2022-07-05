package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"os"
	mock_main "produce-subscription-change/mocks"
	mock_service "produce-subscription-change/mocks/service"
	"produce-subscription-change/service"
	"testing"
)

func TestMainFunction(t *testing.T) {
	_ = os.Setenv("TEST_MAIN", "yes")
	main()
	_ = os.Setenv("TEST_MAIN", "no")
}

func TestProduceWrapFailedDatabaseConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_ = os.Setenv("MYSQL_DSN", "badDatabaseDSN")
	logger := mock_main.NewMockLogger(ctrl)
	logger.EXPECT().Fatalln(gomock.Eq("Failed database connection:"))
	produceWrapper(logger)
}

func TestProduceWrapFailedDataFetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	_ = os.Setenv("MYSQL_DSN", "u:p@/d")
	logger := mock_main.NewMockLogger(ctrl)
	logger.EXPECT().Fatalln(gomock.Eq("Failed data fetch:"))
	produceWrapper(logger)
}

func TestProduceNothingToProcess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	logger := mock_main.NewMockLogger(ctrl)
	logger.EXPECT().Println(gomock.Eq("Nothing to process"))
	database := mock_service.NewMockDatabase(ctrl)
	messages := make([]service.Message, 0)
	database.EXPECT().LoadNewChanges().Return(&messages, nil)
	messaging := mock_service.NewMockMessaging(ctrl)
	produce(logger, database, messaging, func() {
		// function close, nothing to close
	})
}

func TestProduceFailedSendMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	logger := mock_main.NewMockLogger(ctrl)
	logger.EXPECT().Fatalln(gomock.Eq("Failed send messages:"))
	database := mock_service.NewMockDatabase(ctrl)
	messages := make([]service.Message, 0)
	messages = append(messages, service.Message{})
	database.EXPECT().LoadNewChanges().Return(&messages, nil)
	messaging := mock_service.NewMockMessaging(ctrl)
	messaging.EXPECT().SendMessages(gomock.Any(), gomock.Any()).Return(errors.New("some error"))
	produce(logger, database, messaging, func() {
		// function close, nothing to close
	})
}

func TestProduceFailedMarkMessageAsProcessed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	logger := mock_main.NewMockLogger(ctrl)
	logger.EXPECT().Printf(gomock.Eq("Message: %v\n"), gomock.Any()).Times(3)
	logger.EXPECT().Println(gomock.Eq("Message not marked as processed:"), gomock.Any()).Times(3)
	database := mock_service.NewMockDatabase(ctrl)
	messages := make([]service.Message, 0)
	messages = append(messages, service.Message{})
	messages = append(messages, service.Message{})
	messages = append(messages, service.Message{})
	database.EXPECT().LoadNewChanges().Return(&messages, nil)
	database.EXPECT().MarkChangeAsProcessed(gomock.Any()).Return(errors.New("some error")).Times(3)
	messaging := mock_service.NewMockMessaging(ctrl)
	messaging.EXPECT().SendMessages(gomock.Any(), gomock.Any()).DoAndReturn(
		func(messages *[]service.Message, success func(*service.Message)) error {
			for _, message := range *messages {
				success(&message)
			}
			return nil
		},
	)
	produce(logger, database, messaging, func() {
		// function close, nothing to close
	})
}

func TestProduceMarkMessageAsProcessed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	logger := mock_main.NewMockLogger(ctrl)
	logger.EXPECT().Printf(gomock.Eq("Message processed: %v\n"), gomock.Any()).Times(4)
	database := mock_service.NewMockDatabase(ctrl)
	messages := make([]service.Message, 0)
	messages = append(messages, service.Message{})
	messages = append(messages, service.Message{})
	messages = append(messages, service.Message{})
	messages = append(messages, service.Message{})
	database.EXPECT().LoadNewChanges().Return(&messages, nil)
	database.EXPECT().MarkChangeAsProcessed(gomock.Any()).Return(nil).Times(4)
	messaging := mock_service.NewMockMessaging(ctrl)
	messaging.EXPECT().SendMessages(gomock.Any(), gomock.Any()).DoAndReturn(
		func(messages *[]service.Message, success func(*service.Message)) error {
			for _, message := range *messages {
				success(&message)
			}
			return nil
		},
	)
	produce(logger, database, messaging, func() {
		// function close, nothing to close
	})
}
