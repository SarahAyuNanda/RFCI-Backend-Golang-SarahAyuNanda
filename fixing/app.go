package main

import (
	"fmt"
)

func main() {
	collection := []int{3, 5, 2, -4, 8, 11}
	fmt.Println(foo(7, collection))
}

func foo(input int, numbers []int) (result [][]int) {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum := numbers[i] + numbers[j]
			if sum == input {
				pair := []int{numbers[i], numbers[j]}
				result = append(result, pair)
			}
		}
	}
	return
}
