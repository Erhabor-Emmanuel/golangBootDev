package interfaces

import (
	"fmt"
	"time"
)

type birthDayMessage struct {
	birthDayTime  time.Time
	recipientName string
}

type sendingReport struct {
	reportName    string
	numberOfSeats int
}

// Interface
type message interface {
	getMessage() string
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

func (bm birthDayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthDayTime)
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf("Your %s report is ready. You have sent %v total number of seats", sr.reportName, sr.numberOfSeats)
}

// Because the test function takes an interface, we can pass into it any struct that
// implements that interface
func test(m message) {
	sendMessage(m)
	fmt.Println("=====================")
}

func main() {
	test(sendingReport{
		reportName:    "samuel",
		numberOfSeats: 3,
	})
	test(birthDayMessage{
		recipientName: "noel",
		birthDayTime:  time.Date(1994, 07, 21, 0, 0, 0, 0, time.UTC),
	})
}
