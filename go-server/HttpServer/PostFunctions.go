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
	addToken(database.NewDB(), token, post, w)
}

func UserReg(w http.ResponseWriter, r *http.Request) {
	var body = ReadPost(r)
	var post utils.RegJson

	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	if database.NewUserAccount(database.NewDB(), post) {
		var token = utils.RandString(30)
		addToken(database.NewDB(), token, utils.LoginJson{Username: post.Username, Password: post.Password}, w)
	} else {
		sendToken(w, "", false)
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
	fmt.Println("/places/get/30")
	var places = database.GetPlacesMounth(database.NewDB())
	SendJsonArray(w, places)
}

func TicketAdd(w http.ResponseWriter, r *http.Request) {
	body := ReadPost(r)
	var post utils.TicketAdd
	err := json.Unmarshal(body, &post)
	utils.ServerError(err)
	db := database.NewDB()
	userId := database.GetUserIdByToken(db, post.Token)
	res := database.AddTicket(db, utils.Ticket{UserId: fmt.Sprint(userId), PlaceId: post.PlaceId})
	SendJson(w, map[string]string{
		"success": fmt.Sprint(res),
	})
}

func TicketGetMounth(w http.ResponseWriter, r *http.Request) {
	body := ReadPost(r)
	var post string
	err := json.Unmarshal(body, &post)
	utils.ServerError(err)
	var places = database.TicketGetMounth(database.NewDB(), post)
	SendJsonArray(w, places)
}
