package httpserver

import (
	"fmt"
	"net/http"
	"strconv"
)

func StartHttpServer(port int) {
	handler := http.HandlerFunc(HttpHandler)
	fmt.Println("RequestManager is lintening on", port)
	http.ListenAndServe(":"+strconv.Itoa(port), handler)
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

func PostHandler(w http.ResponseWriter, r *http.Request) {

}
