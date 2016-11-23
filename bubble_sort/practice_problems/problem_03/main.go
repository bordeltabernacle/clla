package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	users := []User{
		User{
			FirstName: "Jon",
			LastName:  "Calhoun",
		},
		User{
			FirstName: "Bob",
			LastName:  "Smith",
		},
		User{
			FirstName: "John",
			LastName:  "Smith",
		},
		User{
			FirstName: "Larry",
			LastName:  "Green",
		},
		User{
			FirstName: "Joseph",
			LastName:  "Page",
		},
		User{
			FirstName: "George",
			LastName:  "Costanza",
		},
		User{
			FirstName: "Jerry",
			LastName:  "Seinfeld",
		},
		User{
			FirstName: "Shane",
			LastName:  "Calhoun",
		},
	}

	fmt.Println("Unsorted: ", users)

	bubbleSort(users)
	fmt.Println("Sorted: ", users)
}

func bubbleSort(users []User) {
	N := len(users)
	for i := 0; i < N; i++ {
		if !sweep(users, i) {
			return
		}
	}
}

func sweep(users []User, prevPasses int) bool {
	N := len(users)
	firstIndex := 0
	secondIndex := 1
	didSwap := false

	for secondIndex < (N - prevPasses) {
		firstUser := users[firstIndex]
		secondUser := users[secondIndex]

		if firstUser.LastName > secondUser.LastName {
			users[firstIndex] = secondUser
			users[secondIndex] = firstUser
			didSwap = true
		}

		if firstUser.LastName == secondUser.LastName {
			if firstUser.FirstName > secondUser.FirstName {
				users[firstIndex] = secondUser
				users[secondIndex] = firstUser
				didSwap = true
			}
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}
