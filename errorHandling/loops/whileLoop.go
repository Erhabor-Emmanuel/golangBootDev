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
	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println("numbers: fizz")
		} else if i%5 == 0 {
			fmt.Println("numbers: buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("numbers: fizzbuzz")
		} else {
			fmt.Printf("numbers: %v\n", i)
		}
	}
}

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func testPrime(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===========================")
}

func main() {
	fizzbuzz()
}
