package spe_test

import "fmt"

func narsis() bool {
	var number, tempNumber, remainder int
	var result int = 0
	fmt.Scan(&number)

	tempNumber = number

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber /= 10
		if tempNumber == 0 {
			break
		}
	}

	if result == number {
		fmt.Println("True")
		return true
	} else {
		fmt.Println("False")
		return false
	}
}
