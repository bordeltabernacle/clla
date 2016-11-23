package main

import "fmt"

func main() {
	animals := []string{
		"dog",
		"cat",
		"alligator",
		"Cheetah",
		"rat",
		"moose",
		"cow",
		"mink",
		"porcupine",
		"dung beetle",
		"camel",
		"steer",
		"bat",
		"hamster",
		"horse",
		"colt",
		"bald eagle",
		"frog",
		"rooster",
	}

	fmt.Println("Unsorted: ", animals)

	bubbleSort(animals)
	fmt.Println("Sorted: ", animals)
}

func bubbleSort(animals []string) {
	N := len(animals)
	for i := 0; i < N; i++ {
		if !sweep(animals, i) {
			return
		}
	}
}

func sweep(animals []string, prevPasses int) bool {
	N := len(animals)
	firstIndex := 0
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		firstAnimal := animals[firstIndex]
		secondAnimal := animals[secondIndex]

		if firstAnimal > secondAnimal {
			animals[firstIndex] = secondAnimal
			animals[secondIndex] = firstAnimal
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
