package main

import "fmt"

var (
	Stats = map[string]int{}
)

func CreateUser(user string) {
	fmt.Println("Creating user", user)
	create, ok := Stats["create"]
	if !ok {
		Stats["create"] = 1
		return
	}
	create++
	Stats["create"] = create
}

func UpdateUser(user string) {
	fmt.Println("Updating user", user)

	update, ok := Stats["update"]
	if !ok {
		Stats["update"] = 1
		return
	}
	update++
	Stats["update"] = update
}

func PurgeStats() {
	delete(Stats, "update")
	delete(Stats, "create")
}
