package narsis

import "fmt"

func narsis(number int) bool {
	var tempNumber, remainder int
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
