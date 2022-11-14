package blueocean

import "fmt"

func blueocean(slice1, slice2 []int) []int {

	var result []int
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[0] {
			result = append(result, slice1[i])
		}

	}

	fmt.Println(result)
	return result
}
