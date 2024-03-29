package main

import (
	"fmt"
	"github.com/nickbadlose/mastering-go/postgres"
	"math/rand"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	postgres.Hostname = "localhost"
	postgres.Port = 5432
	postgres.Username = "nickbadlose"
	postgres.Password = "pass"
	postgres.Database = "go"

	users, err := postgres.ListUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(*user)
	}

	username := getString(5)
	u := &postgres.User{
		Username:    username,
		Name:        "Miahlis",
		Surname:     "Tsoukalos",
		Description: "Sensei",
	}
	userID, err := postgres.AddUser(u)
	if err != nil {
		panic(err)
	}

	// this will error adding the same user
	_, err = postgres.AddUser(u)
	if err != nil {
		fmt.Println(err)
	}

	err = postgres.DeleteUSer(userID)
	if err != nil {
		panic(err)
	}

	// this will error as use already deleted
	err = postgres.DeleteUSer(userID)
	if err != nil {
		fmt.Println(err)
	}

	userID, err = postgres.AddUser(u)
	if err != nil {
		panic(err)
	}

	u = &postgres.User{
		Username:    username,
		Name:        "Miahlis",
		Surname:     "Tsoukalos",
		Description: "Updated Sensei",
	}
	err = postgres.UpdateUser(u)
	if err != nil {
		panic(err)
	}

	users, err = postgres.ListUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(*user)
	}
}
