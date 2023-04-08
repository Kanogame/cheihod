package httpserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	database "main/DataBase"
	utils "main/Utils"
	"net/http"
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
	fmt.Fprintf(w, string(jsonRes))
}

func sendToken(w http.ResponseWriter, token string) {
	var data = map[string]string{
		"success": "true",
		"token":   token,
	}
	SendJson(w, data)
}

func addToken(db *sql.DB, token string, name string) {
	data := utils.NewToken{Token: token, Userid: database.GetUserIdByName(db, name)}
	database.NewToken(db, data)
}
