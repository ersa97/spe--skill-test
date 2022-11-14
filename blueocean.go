package spe_test

import "fmt"

func blueocean() []int {
	var length int
	fmt.Scan(&length)

	slice1 := make([]int, length)

	for i := 0; i < length; i++ {
		fmt.Scan(&slice1[i])
	}
	fmt.Println(slice1)
	fmt.Scan(&length)

	slice2 := make([]int, length)

	for i := 0; i < length; i++ {
		fmt.Scan(&slice2[i])
	}

	fmt.Println(slice2)

	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				slice1 = append(slice1[:i], slice1[i+1:]...)
			}
		}
	}

	fmt.Println(slice1)
	return slice1
}
