package httpserver

import (
	"encoding/json"
	"fmt"
	utils "main/Utils"
	"net/http"
)

func SendJson(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	jsonRes, err := json.Marshal(data)
	utils.ServerError(err)
	fmt.Fprintf(w, string(jsonRes))
}

func SendToken(w http.ResponseWriter) {
	var data = map[string]string{
		"success": "true",
		"token":   utils.RandString(30),
	}
	SendJson(w, data)
}
