package create

import (
	"envmanager/pkg/db"
	"envmanager/pkg/general/random"
	"log/slog"
)

var sql_stm[] string = []string { 
	`create table IF NOT EXISTS "User" (userid text PRIMARY KEY, username text UNIQUE, email text, password text)`, 
	`create table IF NOT EXISTS "Service" (service_id text PRIMARY KEY, userid text, service_name text, envs text, foreign key(userid) references "User"(userid))`,
}


func CreateDefaultTable() {
	db := db.Connect()
	defer db.Close()

	for _, sql := range sql_stm {
		_, err := db.Exec(sql)
		if err != nil {
			slog.Error("Error creating table: ", err)
			return
		}
	}
}


func CreateUser(username string, hashed_password string, email string) error {
	db := db.Connect()
	defer db.Close()

	userid := random.MakeUuid()

	sql := `insert into "User" (userid, username, email, password) values ($1, $2, $3, $4)`
	_, err := db.Exec(sql, userid, username, email, hashed_password)
	if err != nil {
		slog.Error("Error inserting user: ", err)
		return err
	}

	return nil
}

func CreateService(userid string, service_id string, service_name string, encrypted_envs string) error {
	db := db.Connect()
	defer db.Close()

	sql := `insert into "Service" (service_id, userid, service_name, envs) values ($1, $2, $3, $4)`
	_, err := db.Exec(sql, service_id, userid, service_name, encrypted_envs)
	if err != nil {
		slog.Error("Error inserting service: ", err)
		return err
	}

	return nil
}