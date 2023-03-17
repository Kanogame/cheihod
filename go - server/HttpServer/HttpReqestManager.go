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
		fmt.Fprintf(w, "nothing to see here")
	} else if r.Method == "POST" {
		PostHandler(w, r)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

}
