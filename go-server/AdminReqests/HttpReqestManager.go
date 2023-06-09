package adminreqests

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func StartHttpServer(port int) {
	var c = cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:10234", "http://192.168.1.7:10234", "http://176.65.35.172:80/admin", "http://176.65.35.172/admin", "http://176.65.35.172"},
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
	"/place/add": AddPlace,
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	PostFunctions[path](w, r)
}
