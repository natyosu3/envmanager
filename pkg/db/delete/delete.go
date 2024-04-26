package delete

import (
	"envmanager/pkg/db"
	"envmanager/pkg/db/common"
	"errors"
	"log/slog"
)


func DeleteService(service_id string, userid string) error {
	db := db.Connect()
	defer db.Close()
	
	owner, err := common.CheckOwner(userid, service_id)
	if err != nil {
		return err
	}
	if !owner {
		return errors.New("you are not the owner of this service")
	}
	result, err := db.Exec(`DELETE FROM "Service" WHERE service_id = $1`, service_id)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	if rowsAffected == 0 {
		return errors.New("service not found")
	} else {
		return nil
	}
}