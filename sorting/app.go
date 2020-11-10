package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{4, 9, 7, 5, 8, 9, 3}
	swap := sorting(numbers)
	fmt.Printf("\nJumlah swap = %d", swap)
}

func sorting(numbers []int) (swapCount int) {
	for i := 0; i < swapCount+1; i++ {
		for index := 1; index < len(numbers); index++ {
			if numbers[index] < numbers[index-1] {
				pairSwap := []int{}
				swapCount++
				store := numbers[index-1]
				numbers[index-1] = numbers[index]
				numbers[index] = store
				pairSwap = append(pairSwap, numbers[index-1], numbers[index])
				resultSwap := arrayToString(numbers)
				fmt.Printf("%d. %v -> %s\n", i+1, pairSwap, resultSwap)
				break
			}
		}
	}
	return
}

func arrayToString(numbers []int) (stringNumbers string) {
	arrayStrings := []string{}
	for i := 0; i < len(numbers); i++ {
		str := strconv.Itoa(numbers[i])
		arrayStrings = append(arrayStrings, str)
		stringNumbers = strings.Join(arrayStrings, " ")
	}
	return
}
