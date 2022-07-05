package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_ToJson(t *testing.T) {
	testTable := []struct {
		message       Message
		expectedJson  string
		expectedError error
	}{
		{
			message: Message{
				SType:     Email,
				Id:        "ivan@antonov.site",
				Value:     true,
				ChangeId:  "32130",
				Confirmed: true,
			},
			expectedJson: `{"type":0,"id":"ivan@antonov.site","value":true,"confirmed":true}`,
		},
		{
			message: Message{
				SType:     Sms,
				Id:        "+7(930)288-08-07",
				Value:     false,
				ChangeId:  "32130",
				Confirmed: false,
			},
			expectedJson: `{"type":1,"id":"+7(930)288-08-07","value":false,"confirmed":false}`,
		},
		{
			message:      Message{},
			expectedJson: `{"type":0,"id":"","value":false,"confirmed":false}`,
		},
	}

	for _, testCase := range testTable {
		data, err := testCase.message.ToJson()
		assert.Nil(t, err)
		assert.Equal(t, testCase.expectedJson, string(data))
	}
}
