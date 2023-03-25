package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "success")
}

func UserReg(w http.ResponseWriter, r *http.Request) {
	var body = readPost(r)
	var post utils.RegJson

	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	fmt.Println(post)

	fmt.Fprintf(w, "success")
}
