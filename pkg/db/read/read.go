package read

import (
	"envmanager/pkg/db"
	"log/slog"
)

var (
	id string
)

func ReadUser(username string) (error) {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query(`SELECT id FROM "User" WHERE name = $1`, username)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
		slog.Info("User found: " + username + id)
	}
	
	return nil
}