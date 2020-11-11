package repository

import "github.com/golang/sarahayunanda/refactory/client-server/api/master/model"

type DataRepo interface {
	CreateData(model.Data) error
}