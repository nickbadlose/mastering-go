package postgres

import (
	_ "github.com/lib/pq"
)

type User struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}
