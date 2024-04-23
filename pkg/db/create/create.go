package create

import (
	"envmanager/pkg/db"
	"fmt"
	"envmanager/pkg/general/random"
)

var sql_stm[] string = []string { 
	`create table IF NOT EXISTS "User" (id text PRIMARY KEY, name text UNIQUE, email text, password text)`, 
	`create table IF NOT EXISTS "element" (id int PRIMARY KEY, name text, value text)`, 
}


func CreateDefaultTable() {
	db := db.Connect()
	defer db.Close()

	for _, sql := range sql_stm {
		_, err := db.Exec(sql)
		if err != nil {
			fmt.Println("Error creating table: ", err)
			return
		}
	}
}


func CreateUser(username string, hashed_password string, email string) error {
	db := db.Connect()
	defer db.Close()

	userid := random.MakeRandomId()

	sql := `insert into "User" (id, name, email, password) values ($1, $2, $3, $4)`
	_, err := db.Exec(sql, userid, username, email, hashed_password)
	if err != nil {
		fmt.Println("Error creating user: ", err)
		return err
	}

	return nil
}