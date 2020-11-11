package master

import (
	"github.com/golang/sarahayunanda/refactory/client-server/api/master/controller"
	"github.com/golang/sarahayunanda/refactory/client-server/api/master/repository"
	"github.com/golang/sarahayunanda/refactory/client-server/api/master/usecase"
	"github.com/gorilla/mux"
)


func Init(r *mux.Router) {
	dataRepo := repository.InitDataRepoImple()
	dataUsecase := usecase.InitDataUsecase(dataRepo)
	controller.DataController(r, dataUsecase)
}