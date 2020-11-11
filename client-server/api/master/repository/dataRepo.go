package repository

import (
	"net/http"

	"github.com/golang/sarahayunanda/refactory/client-server/api/master/model")


type DataRepo interface {
	CreateData(model.Data, *http.Request) error
}