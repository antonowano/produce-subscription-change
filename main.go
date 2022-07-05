package main

import (
	"log"
	"os"
	"produce-subscription-change/service"
	"strings"
)

//go:generate mockgen -source=main.go -destination=mocks/main.go

type Logger interface {
	Fatalln(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

func main() {
	produceWrapper(log.Default())
}

func produceWrapper(log Logger) {
	var (
		dsn       = os.Getenv("MYSQL_DSN")
		addresses = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
		username  = os.Getenv("KAFKA_USERNAME")
		password  = os.Getenv("KAFKA_PASSWORD")
		testMain  = os.Getenv("TEST_MAIN")
	)

	// break execution for test main function
	if testMain == "yes" {
		return
	}

	db, err := service.NewDB(dsn)
	if err != nil {
		log.Fatalln("Failed database connection:", err)
		return
	}

	database := service.NewDatabase(db)
	writer := service.NewWriter(addresses, username, password)
	messaging := service.NewMessaging(writer)

	produce(
		log,
		database,
		messaging,
		func() {
			_ = writer.Close()
			_ = database.Close()
		},
	)
}

func produce(log Logger, database service.Database, messaging service.Messaging, close func()) {
	messages, err := database.LoadNewChanges()
	if err != nil {
		close()
		log.Fatalln("Failed data fetch:", err)
		return
	}

	if len(*messages) == 0 {
		close()
		log.Println("Nothing to process")
		return
	}

	err = messaging.SendMessages(messages, func(message *service.Message) {
		err2 := database.MarkChangeAsProcessed(message)
		if err2 != nil {
			log.Printf("Message: %v\n", message)
			log.Println("Message not marked as processed:", err)
		} else {
			log.Printf("Message processed: %v\n", message)
		}
	})

	if err != nil {
		close()
		log.Fatalln("Failed send messages:", err)
	}
}
