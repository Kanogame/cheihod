package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func StartHttpServer(port int) {
	var c = cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5501", "http://192.168.1.7:5501", "http://176.65.35.172:80", "http://176.65.35.172:80/api", "http://176.65.35.172", "http://176.65.35.172/api"},
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
	"/api/login":         UserLogin,
	"/api/reg":           UserReg,
	"/api/token":         TokenName,
	"/api/token/full":    TokenFull,
	"/api/places/get/30": PlacesGetMounth,
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	PostFunctions[path](w, r)
}
