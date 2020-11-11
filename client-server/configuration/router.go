package configuration

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() (router *mux.Router) {
	router = mux.NewRouter()
	return
}

func RunServer(router *mux.Router) {
	port := "8080"
	log.Println("Starting web server at port", port)
	err := http.ListenAndServe("localhost:"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
