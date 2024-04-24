package read

import (
	db_model "envmanager/pkg/db/model"
	"envmanager/pkg/model"
	"envmanager/pkg/db"
)


func ReadUser(username string) (*db_model.User, error) {
	db := db.Connect()
	defer db.Close()

	var db_user db_model.User

	err := db.QueryRow(`SELECT userid, password FROM "User" WHERE username = $1`, username).Scan(&db_user.Userid, &db_user.Password)
	if err != nil {
		return nil, err
	}
	
	return &db_user, nil
}

func ReadService(userid string) ([]model.Service_model, error) {
	db := db.Connect()
	defer db.Close()

	var service[] model.Service_model

	rows, err := db.Query(`SELECT service_name, service_id FROM "Service" WHERE userid = $1`, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			service_name string
			service_id string
		)
		err := rows.Scan(&service_name, &service_id)
		if err != nil {
			return nil, err
		}
		service = append(service, model.Service_model{Service_name: service_name, Service_id: service_id})
	}

	return service, nil
}


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