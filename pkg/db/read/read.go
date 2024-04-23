package read

import (
	"envmanager/pkg/db/model"
	"envmanager/pkg/db"
)


func ReadUser(username string) (*model.User, error) {
	db := db.Connect()
	defer db.Close()

	var db_user model.User

	err := db.QueryRow(`SELECT userid, password FROM "User" WHERE username = $1`, username).Scan(&db_user.Userid, &db_user.Password)
	if err != nil {
		return nil, err
	}
	
	return &db_user, nil
}