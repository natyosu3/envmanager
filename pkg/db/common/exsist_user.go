package common

import (
	"envmanager/pkg/db"
	"log/slog"
)


func ExsistUser(username string) (bool) {
	db := db.Connect()
	defer db.Close()

	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM "User" WHERE name = $1`, username).Scan(&count)
	if err != nil {
		slog.Error("Error checking if user exsist: " + err.Error())
		return false
	}
	return count > 0
}