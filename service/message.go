package service

import (
	"encoding/json"
)

//go:generate mockgen -source=message.go -destination=../mocks/service/message.go

type SubscriptionType int8

const (
	Email SubscriptionType = 0
	Sms                    = 1
)

type Message struct {
	SType     SubscriptionType `json:"type"`
	Id        string           `json:"id"`
	Value     bool             `json:"value"`
	Confirmed bool             `json:"confirmed"`
	ChangeId  string           `json:"-"`
}

func (m *Message) ToJson() ([]byte, error) {
	return json.Marshal(m)
}
