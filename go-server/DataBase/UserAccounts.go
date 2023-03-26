package database

import (
	"database/sql"
	"fmt"
	utils "main/Utils"
)

func NewUserAccount(db *sql.DB, data utils.RegJson) bool {
	res, err := db.Exec(fmt.Sprintf("INSERT INTO Users (username, passwrd, email) VALUES('%v', '%v', '%v')", data.Username, data.Password, data.Email))
	success := utils.UserError(err)
	if res == nil {
		utils.UserError(fmt.Errorf("fuck"))
	}
	return success
}

func VerifyUserAccount(db *sql.DB, data utils.LoginJson) bool {
	res, err := db.Query("SELECT passwrd FROM Users WHERE usermane = ?", data.Username)
	utils.UserError(err)

	var password string

	for res.Next() {
		err := res.Scan(&password)
		if !utils.UserError(err) {
			return false
		}
	}

	return password == data.Password
}
