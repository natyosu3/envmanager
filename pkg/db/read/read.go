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


// return service_name, encrypted_envs, error
func ReadServiceDetail(service_id string) (string, string, error) {
	db := db.Connect()
	defer db.Close()

	var service_name string
	var encrypted_envs string

	err := db.QueryRow(`SELECT service_name, envs FROM "Service" WHERE service_id = $1`, service_id).Scan(&service_name, &encrypted_envs)
	if err != nil {
		return "", "", err
	}

	return service_name, encrypted_envs, nil
}
