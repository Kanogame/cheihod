package httpserver

import (
	"encoding/json"
	"fmt"
	database "main/DataBase"
	utils "main/Utils"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var body = readPost(r)
	var post utils.LoginJson

	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	db := database.NewDB()

	var token = utils.RandString(30)
	addToken(db, token, post.Username)
	sendToken(w, token)
}

func UserReg(w http.ResponseWriter, r *http.Request) {
	var body = readPost(r)
	var post utils.RegJson

	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	db := database.NewDB()
	if database.NewUserAccount(db, post) {
		fmt.Fprint(w, "success")
		fmt.Println("success")
	} else {
		fmt.Fprint(w, "exist")
	}

}
