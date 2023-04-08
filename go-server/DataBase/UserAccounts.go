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

func GetUserIdByToken(db *sql.DB, token string) int {
	res, err := db.Query("SELECT userid FROM Token WHERE token = ?", token)
	utils.UserError(err)

	var id int

	for res.Next() {
		err := res.Scan(&id)
		utils.UserError(err)
	}
	return id
}

func GetUserById(db *sql.DB, id int) utils.RegJson {
	res, err := db.Query("SELECT username, passwrd, email FROM Users WHERE id = ?", id)
	utils.UserError(err)

	var user utils.RegJson

	for res.Next() {
		err := res.Scan(&user.Username, &user.Password, &user.Email)
		utils.UserError(err)
	}
	return user
}

func GetNameByToken(db *sql.DB, token string) string {
	return GetUserById(db, GetUserIdByToken(db, token)).Username
}

func GetUserByToken(db *sql.DB, token string) utils.RegJson {
	return GetUserById(db, GetUserIdByToken(db, token))
}
