package main

import "fmt"

func main() {
	numbers := []string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}
	fmt.Println("Unsorted: ", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted: ", numbers)
}

func bubbleSort(numbers []string) {
	N := len(numbers)
	for i := 0; i < N; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []string, prevPasses int) bool {
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
