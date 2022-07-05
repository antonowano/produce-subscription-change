package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type BadRows struct {
	Rows
}

func (rows *BadRows) Next() bool {
	return true
}

func (rows *BadRows) Scan(_ ...interface{}) error {
	return errors.New("some error")
}

type BadDB struct {
	DBImpl
}

func (db *BadDB) Query(_ string, _ ...interface{}) (Rows, error) {
	return &BadRows{}, nil
}

var dsn = os.Getenv("MYSQL_DSN")

func TestNewDB(t *testing.T) {
	_, err := NewDB("u0:p0@/dbname")
	assert.Nil(t, err)
}

func TestNewDB2(t *testing.T) {
	_, err := NewDB("sdfsdf")
	assert.Error(t, err)
}

func TestNewDatabase(t *testing.T) {
	db, err := NewDB(dsn)
	assert.Nil(t, err)
	NewDatabase(db)
}

func TestDatabase_Close(t *testing.T) {
	db, err := NewDB(dsn)
	assert.Nil(t, err)
	database := NewDatabase(db)
	err = database.Close()
	assert.Nil(t, err)
}

func TestDatabase_LoadNewChanges(t *testing.T) {
	db, err := NewDB(dsn)
	assert.Nil(t, err)
	database := NewDatabase(db)
	messages, err := database.LoadNewChanges()
	assert.Nil(t, err)
	assert.Equal(t, 12, len(*messages))
	assert.Equal(t, SubscriptionType(1), (*messages)[0].SType)
	assert.Equal(t, "+7(123)123-12-31", (*messages)[0].Id)
	assert.Equal(t, "1", (*messages)[0].ChangeId)
	assert.Equal(t, true, (*messages)[0].Confirmed)
	assert.Equal(t, false, (*messages)[0].Value)
	err = database.Close()
	assert.Nil(t, err)
}

func TestDatabase_LoadNewChangesBadAuth(t *testing.T) {
	db, err := NewDB("u0:p0@/dbname")
	assert.Nil(t, err)
	database := NewDatabase(db)
	_, err = database.LoadNewChanges()
	assert.Error(t, err)
}

func TestDatabase_LoadNewChangesBadScan(t *testing.T) {
	database := NewDatabase(&BadDB{})
	_, err := database.LoadNewChanges()
	assert.Error(t, err)
}

func TestDatabase_MarkChangeAsProcessed(t *testing.T) {
	db, err := NewDB(dsn)
	assert.Nil(t, err)
	database := NewDatabase(db)
	err = database.MarkChangeAsProcessed(&Message{
		ChangeId: "2",
	})
	assert.Nil(t, err)
}
