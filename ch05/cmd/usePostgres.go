package main

import (
	"fmt"
	"github.com/nickbadlose/mastering-go/postgres"
)

func main() {
	postgres.Hostname = "localhost"
	fmt.Println(postgres.Port)
	fmt.Println(postgres.Hostname)
}
