package models

import (
	"database/sql"

	"github.com/peter/go-auth/config"
	"github.com/peter/go-auth/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConfig()
	if err != nil {
		panic(err)
	}
	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
	row, err := u.db.Query("SELECT * FROM users WHERE "+fieldName+" = ?", fieldValue)
	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&user.ID, &user.Username, &user.Password)
	}
	return nil
}
