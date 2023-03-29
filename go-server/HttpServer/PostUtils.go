package httpserver

import (
	"encoding/json"
	"fmt"
	utils "main/Utils"
	"net/http"
)

func SendJson(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	var Res = map[string]string{
		"success": "true",
		"token":   "testToken",
	}
	jsonRes, err := json.Marshal(data)
	utils.ServerError(err)
	fmt.Fprintf(w, string(jsonRes))
}
