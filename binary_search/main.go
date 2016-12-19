package main

import "fmt"

func main() {
	t := 6
	l := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	fmt.Println("Looking for", t, "in the sorted list:", l)

	index := binarySearch(l, t)
	if index >= 0 {
		fmt.Println("Found the number", t, "at:", index)
	} else {
		fmt.Println("Didn't find the number", t, ":(")
	}
}

func binarySearch(l []int, t int) int {
	low := 0
	high := len(l) - 1

	for low <= high {
		mid := low + (high-low)/2
		midValue := l[mid]
		fmt.Println("Middle value is:", midValue)

		if midValue == t {
			return mid
		} else if midValue > t {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
