package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	database "main/DataBase"
	utils "main/Utils"
	"net/http"
)

func readPost(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	utils.ServerError(err)
	return body
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var body = readPost(r)
	var post utils.LoginJson

	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	fmt.Println(post)

	w.Header().Set("Content-Type", "application/json")
	var Res = map[string]string{
		"success": "true",
		"token":   "testToken",
	}
	jsonRes, err := json.Marshal(Res)
	utils.ServerError(err)
	fmt.Fprintf(w, string(jsonRes))
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
