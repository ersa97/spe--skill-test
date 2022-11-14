package spe_test

import "fmt"

func outliers(slice []int) any {
	var odd, even, outlier, result int

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		outlier = 1
	} else if odd > even {
		outlier = 0
	}

	if odd == 0 || even == 0 {
		return false
	}

	fmt.Println("outlier", outlier)

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == outlier {
			result = slice[i]
		}
	}
	fmt.Println(result)
	return result
}
