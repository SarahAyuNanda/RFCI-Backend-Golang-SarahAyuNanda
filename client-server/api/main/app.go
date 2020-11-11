package main

import (
	"github.com/golang/sarahayunanda/refactory/client-server/api/master"
	"github.com/golang/sarahayunanda/refactory/client-server/configuration")


func main() {
	router := configuration.CreateRouter()
	master.Init(router)
	configuration.RunServer(router)
}