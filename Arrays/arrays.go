package Arrays

import (
	"fmt"
)

func getMessageWithRetries() [6]string {
	return [6]string{
		"Click here to sign up",
		"Pretty please click here",
		"We beg you to sign up",
	}
}

func send(name string, doneAt int) {
	fmt.Printf("sending to %v....", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {

	}
}
