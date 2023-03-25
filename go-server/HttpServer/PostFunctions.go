package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	utils "main/Utils"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var post utils.LoginJson

	err = json.Unmarshal(body, &post)
	if err != nil {
		panic(err)
	}

	fmt.Println(post)

	fmt.Fprintf(w, "success")
}
