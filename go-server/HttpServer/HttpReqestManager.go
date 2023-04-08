package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func StartHttpServer(port int) {
	var c = cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5501"},
	})

	handler := http.HandlerFunc(HttpHandler)
	fmt.Println("RequestManager is lintening on", port)
	http.ListenAndServe(":"+strconv.Itoa(port), c.Handler(handler))
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := r.URL.Path

		if path == "/" {
			path = "./static/"
		} else {
			path = "./static/" + path
		}

		http.ServeFile(w, r, path)
	} else if r.Method == "POST" {
		PostHandler(w, r)
	}
}

var PostFunctions = map[string]func(http.ResponseWriter, *http.Request){
	"/login":      UserLogin,
	"/reg":        UserReg,
	"/token":      TokenName,
	"/token/full": TokenFull,
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	PostFunctions[path](w, r)
}
