package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//go:generate mockgen -source=database.go -destination=../mocks/service/database.go

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
}

type Result interface {
}

type DB interface {
	Close() error
	Query(query string, args ...interface{}) (Rows, error)
	Exec(query string, args ...interface{}) (Result, error)
}

type Database interface {
	Close() error
	MarkChangeAsProcessed(message *Message) error
	LoadNewChanges() (*[]Message, error)
}

func NewDB(dsn string) (DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	myDB := &DBImpl{
		db: db,
	}

	return myDB, nil
}

type DBImpl struct {
	db *sql.DB
}

func (db *DBImpl) Close() error {
	return db.db.Close()
}

func (db *DBImpl) Query(query string, args ...interface{}) (Rows, error) {
	return db.db.Query(query, args...)
}

func (db *DBImpl) Exec(query string, args ...interface{}) (Result, error) {
	return db.db.Exec(query, args...)
}

func NewDatabase(db DB) Database {
	return &DatabaseImpl{
		db: db,
	}
}

type DatabaseImpl struct {
	db DB
}

func (service *DatabaseImpl) Close() error {
	return service.db.Close()
}

func (service *DatabaseImpl) MarkChangeAsProcessed(message *Message) error {
	sqlQuery := "UPDATE subscription_change SET is_processed = 1 WHERE id = ?"
	_, err := service.db.Exec(sqlQuery, message.ChangeId)
	return err
}

func (service *DatabaseImpl) LoadNewChanges() (*[]Message, error) {
	messages := make([]Message, 0)
	rows, err := service.db.Query("SELECT id, email, email_confirmed, phone, subscription_type, subscription_action" +
		" FROM subscription_change WHERE is_processed = 0 AND `source` != 4")
	if err != nil {
		return nil, err
	}

	var (
		id               string
		email            sql.NullString
		confirmed        bool
		phone            sql.NullString
		subscriptionType SubscriptionType
		action           bool
	)

	for rows.Next() {
		if err := rows.Scan(&id, &email, &confirmed, &phone, &subscriptionType, &action); err != nil {
			return nil, err
		}

		message := Message{
			SType:    subscriptionType,
			Value:    action,
			ChangeId: id,
		}

		switch subscriptionType {
		case Email:
			message.Id = email.String
			message.Confirmed = confirmed
		case Sms:
			message.Id = phone.String
			message.Confirmed = true
		}

		messages = append(messages, message)
	}

	return &messages, nil
}
