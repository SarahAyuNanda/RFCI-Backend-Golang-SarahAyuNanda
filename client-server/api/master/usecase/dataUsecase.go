package usecase

import "github.com/golang/sarahayunanda/refactory/client-server/api/master/model"

type DataUsecase interface {
	PostData(model.Data) error
}