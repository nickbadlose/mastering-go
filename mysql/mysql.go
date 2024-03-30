package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// Connection details
var (
	Hostname = "127.0.0.1"
	Port     = 3306
	Username = ""
	Password = ""
	Database = ""
)

type User struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

func openConnection() (*sql.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@%s:%s/%s",
		Username,
		Password,
		Hostname,
		Port,
		Database,
	)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func exists(username string) (int, error) {
	username = strings.ToLower(username)

	db, err := openConnection()
	if err != nil {
		return 0, err
	}

	rows, err := db.Query(`SELECT "id" FROM "users" WHERE username = $1`, username)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	userID := 0
	for rows.Next() {
		sErr := rows.Scan(&userID)
		if sErr != nil {
			return 0, sErr
		}
	}

	if rows.Err() != nil {
		return 0, rows.Err()
	}

	return userID, nil
}

func AddUser(d *User) (int, error) {
	if d == nil {
		return 0, errors.New("no user provided")
	}

	d.Username = strings.ToLower(d.Username)

	db, err := openConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	userID, err := exists(d.Username)
	if err != nil {
		return 0, err
	}

	insertStatement := `insert into "users" ("username") values ($1)`
	_, err = db.Exec(insertStatement, d.Username)
	if err != nil {
		return 0, err
	}

	userID, err = exists(d.Username)
	if err != nil {
		return 0, err
	}

	insertStatement = `insert into "userdata" ("userid", "name", "surname", "description") values ($1, $2, $3, $4)`
	_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func DeleteUSer(id int) error {
	if id <= 0 {
		return errors.New("id must be a positive, non zero integer")
	}

	db, err := openConnection()
	if err != nil {
		return err
	}

	rows, err := db.Query(`SELECT "username" FROM "users" WHERE id = $1`, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var username string
	for rows.Next() {
		sErr := rows.Scan(&username)
		if sErr != nil {
			return sErr
		}
	}

	if rows.Err() != nil {
		return rows.Err()
	}

	// no user exists for the given ID
	if username == "" {
		return errors.New("user does not exist")
	}

	_, err = db.Exec(`DELETE FROM "userdata" WHERE userid = $1`, id)
	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM "users" WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func ListUsers() ([]*User, error) {
	db, err := openConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT 
    "id","username","name","surname","description" 
	FROM "users","userdata" WHERE users.id = userdata.userid`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0, 1)
	for rows.Next() {
		u := &User{}
		sErr := rows.Scan(&u.ID, &u.Username, &u.Name, &u.Surname, &u.Description)
		if sErr != nil {
			return nil, sErr
		}

		users = append(users, u)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func UpdateUser(u *User) error {
	if u == nil {
		return errors.New("no user provided")
	}

	db, err := openConnection()
	if err != nil {
		return err
	}

	userID, err := exists(u.Username)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`UPDATE "userdata" SET "name"=$1, "surname"=$2, "description"=$3 WHERE "userid"=$4`,
		u.Name,
		u.Surname,
		u.Description,
		userID,
	)

	return err
}
