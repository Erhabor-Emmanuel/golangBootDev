package main

import (
	"fmt"
)

func main() {
	age := 20
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvent(age)

	fmt.Printf("mano firstNames are %v, %v, %v", yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)

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
