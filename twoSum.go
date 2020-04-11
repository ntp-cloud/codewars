package main

import f "fmt"

/**
* the TwoSum function is imported for be used in golang testing
 */
func TwoSum(numbers []int, target int) [2]int {

	// #var.#1
	// 	for i := 0; i < len(numbers); i++ {
	// 		for j := i + 1; j < len(numbers); j++ {
	// 			if numbers[i]+numbers[j] == target {
	// 				return [2]int{i, j}
	// 			}
	// 		}
	// 	}
	// 	return [2]int{0, 0}

	// #var.#2
	for i, vI := range numbers {
		for j, vJ := range numbers {
			if i != j && vI+vJ == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}

func main() {
	f.Println(TwoSum([]int{1, 2, 3}, 4))
}
