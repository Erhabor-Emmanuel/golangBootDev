package loops

import "fmt"

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
	}
	return maxMessagesToSend
}

func fizzbuzz() {
	for i := 0; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println("numbers: fizz")
		}
		if i%5 == 0 {
			fmt.Println("numbers: buzz")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("numbers: fizzbuzz")
		} else {
			fmt.Printf("numbers: %v\n", i)
		}
	}
}

func main() {
	fizzbuzz()
}
