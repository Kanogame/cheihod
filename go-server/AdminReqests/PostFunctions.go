package adminreqests

import (
	"encoding/json"
	database "main/DataBase"
	httpserver "main/HttpServer"
	utils "main/Utils"
	"net/http"
)

func AddPlace(w http.ResponseWriter, r *http.Request) {
	var body = httpserver.ReadPost(r)
	var post utils.AddPlace
	err := json.Unmarshal(body, &post)
	utils.ServerError(err)

	db := database.NewDB()
	var name = database.AddPlace(db, post)
	if name {
		httpserver.SendJson(w, map[string]string{
			"success": "true",
		})
	}
}
