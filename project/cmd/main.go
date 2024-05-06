package main

import (
	"log"
	"net/http"
	"project/web"
)

func main() {
	mux := web.StartServer()
	log.Fatal(http.ListenAndServe(":3000", mux))
}
