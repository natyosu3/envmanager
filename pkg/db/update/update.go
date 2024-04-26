package update

import (
	"envmanager/pkg/db"
)


func UpdateService(service_id string, envs string, userid string) error {
	db := db.Connect()
	defer db.Close()

	_, err := db.Exec(`UPDATE "Service" SET envs = $1 WHERE service_id = $2`, envs, service_id)
	if err != nil {
		return err
	}
	return nil
}