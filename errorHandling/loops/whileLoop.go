package loops

func getMaxMessagesToSend(costMultiplier float64, maxCostPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
}
