package usecase

import (
	"net/http"

	"github.com/golang/sarahayunanda/refactory/client-server/api/master/model")


type DataUsecase interface {
	PostData(model.Data, *http.Request) error
}