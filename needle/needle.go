package needle

import "fmt"

func needle(slice []string, text string) int {
	var result int

	for i := 0; i < len(slice); i++ {
		if slice[i] == text {
			result = i
		}
	}

	fmt.Println(result)
	return result
}
