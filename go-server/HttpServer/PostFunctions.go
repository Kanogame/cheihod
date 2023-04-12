package httpserver

import (
	"encoding/json"
	"fmt"
	database "main/DataBase"
	utils "main/Utils"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var body = ReadPost(r)
	var post utils.LoginJson

	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	var token = utils.RandString(30)
	addToken(database.NewDB(), token, post.Username)
	sendToken(w, token)
}

func UserReg(w http.ResponseWriter, r *http.Request) {
	var body = ReadPost(r)
	var post utils.RegJson

	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	if database.NewUserAccount(database.NewDB(), post) {
		fmt.Fprint(w, "success")
		fmt.Println("success")
	} else {
		fmt.Fprint(w, "exist")
	}
}

func TokenName(w http.ResponseWriter, r *http.Request) {
	var body = ReadPost(r)
	var post string
	err := json.Unmarshal(body, &post)
	utils.ServerError(err)
	var name = database.GetNameByToken(database.NewDB(), post)
	if name != "" {
		SendJson(w, map[string]string{
			"success": "true",
			"name":    name,
		})
	}
}

func TokenFull(w http.ResponseWriter, r *http.Request) {
	var body = ReadPost(r)
	var post string
	err := json.Unmarshal(body, &post)
	utils.ServerError(err)
	var user = database.GetUserByToken(database.NewDB(), post)
	SendJson(w, map[string]string{
		"success": "true",
		"name":    user.Username,
		"email":   user.Email,
	})
}

func PlacesGetMounth(w http.ResponseWriter, r *http.Request) {
	var places = database.GetPlacesMounth(database.NewDB())
	SendJsonArray(w, places)
}
