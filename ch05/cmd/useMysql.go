package main

import (
	"fmt"
	"github.com/nickbadlose/mastering-go/mysql"
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
	mysql.Hostname = "localhost"
	mysql.Port = 3306
	mysql.Username = "nickbadlose"
	mysql.Password = "pass"
	mysql.Database = "go"

	users, err := mysql.ListUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(*user)
	}

	username := getString(5)
	u := &mysql.User{
		Username:    username,
		Name:        "Miahlis",
		Surname:     "Tsoukalos",
		Description: "Sensei",
	}
	userID, err := mysql.AddUser(u)
	if err != nil {
		panic(err)
	}

	// this will error adding the same user
	_, err = mysql.AddUser(u)
	if err != nil {
		fmt.Println(err)
	}

	err = mysql.DeleteUSer(userID)
	if err != nil {
		panic(err)
	}

	// this will error as use already deleted
	err = mysql.DeleteUSer(userID)
	if err != nil {
		fmt.Println(err)
	}

	userID, err = mysql.AddUser(u)
	if err != nil {
		panic(err)
	}

	u = &mysql.User{
		Username:    username,
		Name:        "Miahlis",
		Surname:     "Tsoukalos",
		Description: "Updated Sensei",
	}
	err = mysql.UpdateUser(u)
	if err != nil {
		panic(err)
	}

	users, err = mysql.ListUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(*user)
	}
}
