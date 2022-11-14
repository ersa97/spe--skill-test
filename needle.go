package spe_test

import "fmt"

func needle() int {
	var length, result int
	var text string
	fmt.Scan(&length)

	slice := make([]string, length)

	for i := 0; i < length; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Println(slice)

	fmt.Scan(&text)

	for i := 0; i < length; i++ {
		if slice[i] == text {
			result = i
		}
	}

	fmt.Println(result)
	return result
}
