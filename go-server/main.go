package main

import (
	adminreqests "main/AdminReqests"
	httpserver "main/HttpServer"
)

func main() {
	go httpserver.StartHttpServer(10234)
	adminreqests.StartHttpServer(10235)
}
