package main

import (
	"log"
	"net/http"
	"os"
	"proxy_microservice/handler"
	"proxy_microservice/server"
)

var (
	port = "8081"
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
}

func main() {
	server := server.Server{
		Router: http.NewServeMux(),
	}
	h := handler.Handler{}
	server.InitRoute(&h)
	log.Fatal(http.ListenAndServe(`:`+port, server.Router))

}
