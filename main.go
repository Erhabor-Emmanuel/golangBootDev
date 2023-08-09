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
type messageToSend struct {
	message     string
	phoneNumber int
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
