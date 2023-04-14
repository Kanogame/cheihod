package httpserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	database "main/DataBase"
	utils "main/Utils"
	"net/http"
	"strconv"
)

func ReadPost(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	utils.ServerError(err)
	return body
}

func SendJson(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	jsonRes, err := json.Marshal(data)
	utils.ServerError(err)
	fmt.Fprint(w, string(jsonRes))
}

func SendJsonArray(w http.ResponseWriter, data []map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	jsonRes, err := json.Marshal(data)
	utils.ServerError(err)
	fmt.Fprint(w, string(jsonRes))
}

func sendToken(w http.ResponseWriter, token string, success bool) {
	var data = map[string]string{
		"success": strconv.FormatBool(success),
		"token":   token,
	}
	SendJson(w, data)
}

func addToken(db *sql.DB, token string, post utils.LoginJson, w http.ResponseWriter) {
	if !database.VerifyUserAccount(db, post) {
		sendToken(w, "", false)
		return
	}
	data := utils.NewToken{Token: token, Userid: database.GetUserIdByName(db, post.Username)}
	database.NewToken(db, data)
	sendToken(w, token, true)
}
