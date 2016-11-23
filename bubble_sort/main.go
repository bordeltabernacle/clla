package main

import "fmt"

func main() {
	numbers := []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Unsorted: ", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted: ", numbers)
}

func bubbleSort(numbers []int) {
	N := len(numbers)
	for i := 0; i < N; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int) bool {
	N := len(numbers)
	firstIndex := 0
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		firstNumber := numbers[firstIndex]
		secondNumber := numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
