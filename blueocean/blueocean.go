package blueocean

import "fmt"

func blueocean(slice1, slice2 []int) []int {
	var length int
	fmt.Scan(&length)

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
