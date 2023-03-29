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
	res, err := db.Query("SELECT passwrd FROM Users WHERE username = ?", data.Username)
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

func GetUserIdByName(db *sql.DB, name string) int {
	res, err := db.Query("SELECT id FROM Users WHERE username = ?", name)
	utils.UserError(err)

	var id int

	for res.Next() {
		err := res.Scan(&id)
		utils.UserError(err)
	}
	return id
}

func NewToken(db *sql.DB, data utils.NewToken) bool {
	res, err := db.Exec(fmt.Sprintf("INSERT INTO Token (token, userid) VALUES('%v', '%v')", data.Token, data.Userid))
	success := utils.UserError(err)
	if res == nil {
		utils.UserError(fmt.Errorf("fuck"))
	}
	return success
}
