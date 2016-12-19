package main

import (
	"fmt"
	"strings"
)

// Given a list of sorted words (strings with no spaces),
// search for a user provided word in the list without
// being case sensitive. Return -1 if the word isn't found
// and return the index of the word if it is found.
func main() {
	words := []string{
		"ALLigator",
		"bat",
		"bEEtle",
		"camel",
		"cat",
		"cheetah",
		"COLT",
		"cow",
		"dog",
		"eagle",
		"froG",
		"hamster",
		"horse",
		"mink",
		"moose",
		"porcupine",
		"RaT",
		"rooster",
		"steer",
	}
	fmt.Println("Sorted Words:", words)
	var toFind string
	fmt.Println("What word should we search for? No spaces please!")
	fmt.Scanf("%s", &toFind)
	index := binarySearch(words, toFind)
	if index < 0 {
		fmt.Println("The word", toFind, "could not be found!")
	} else {
		fmt.Println("The word", toFind, "was found at index:", index, words[index])
	}
}

func binarySearch(l []string, t string) int {
	t = strings.ToLower(t)
	low := 0
	high := len(l) - 1

	for low <= high {
		mid := low + (high-low)/2
		v := strings.ToLower(l[mid])
		fmt.Println("Middle value is:", v)

		if v == t {
			return mid
		} else if v > t {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
