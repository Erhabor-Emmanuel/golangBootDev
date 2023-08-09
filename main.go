package main

import (
	"fmt"
	// "golang.org/x/text/message"
)

func main() {
	test(messageToSend{
		message:     "Thanks for signing up",
		phoneNumber: 9080809790,
	})

	test(messageToSend{
		message:     "Thanks for logging in",
		phoneNumber: 9080808790,
	})

	test(messageToSend{
		message:     "Thanks for loging out",
		phoneNumber: 9080879790,
	})

}

// struct keys can hold any type
// The fields of a struct can be accessed using the dot . operator
// myNo = messageToSend{}
// myNo.phoneNumber
type messageToSend struct {
	message     string
	phoneNumber int
	sender      user
	recipient   user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func test(m messageToSend) {
	fmt.Printf("Sending mesage: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("===============================")
}

// Naked Return
func yearsUntilEvent(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}
