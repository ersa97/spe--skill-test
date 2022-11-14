package spe_test

import "fmt"

func outliers() int {
	var length, odd, even, outlier, result int
	fmt.Scan(&length)

	slice := make([]int, length)

	for i := 0; i < length; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < length; i++ {
		if slice[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println(slice)

	if even > odd {
		outlier = 1
	} else {
		outlier = 0
	}

	fmt.Println("outlier", outlier)

	for i := 0; i < length; i++ {
		if slice[i]%2 == outlier {
			result = slice[i]
		}
	}
	fmt.Println(result)
	return result
}
