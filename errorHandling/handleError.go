package golangbootdev

import "fmt"

func sendSmsToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCustomer, err := sendSms(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse, errr := sendSms(msgToSpouse)
	if errr != nil {
		return 0.0, errr
	}
	return costForCustomer + costForSpouse, nil
}

func sendSms(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send text over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}
