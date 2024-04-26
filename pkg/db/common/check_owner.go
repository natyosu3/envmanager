package common

import (
	"envmanager/pkg/db"
)

func CheckOwner(userid string, service_id string) (bool, error) {
	db := db.Connect()
	defer db.Close()

	var count int

	err := db.QueryRow(`SELECT COUNT(*) FROM "Service" WHERE service_id = $1 AND userid = $2`, service_id, userid).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}