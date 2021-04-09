package vatansms

import (
	"log"
	"testing"
)

func TestOneToN_Send1N(t *testing.T) {
	var numberAndMessages []NumberAndMessage
	numberAndMessages = append(numberAndMessages, NumberAndMessage{
		Number:  PhoneVerify("5442666417"),
		Message: "Hi 5442666417",
	})

	smsNN := NToN{}
	smsNN.UserID = 1234
	smsNN.Username = "1234"
	smsNN.Password = "1234"
	smsNN.Sender = "TEST SENDER"
	smsNN.NumberAndMessages = numberAndMessages
	smsNN.Type = "normal"
	res, err := smsNN.SendNN()

	if err != nil {
		t.Error(err)
	}

	log.Println(res)
}
