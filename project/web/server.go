package web

import (
	"log"
	"net/http"
	"project/web/middlewares"
)

func StartServer() *http.ServeMux {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()

	InitRoutes(mux, manager)
	log.Println("Server started on :3000")
	return mux
}
