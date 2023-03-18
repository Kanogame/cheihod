package main

import (
	httpserver "main/HttpServer"
)

func main() {
	httpserver.StartHttpServer(10234)
}
