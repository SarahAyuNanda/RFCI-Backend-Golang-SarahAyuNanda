package usecase

import (
	"log"

	"github.com/golang/sarahayunanda/refactory/client-server/api/master/model"
	"github.com/golang/sarahayunanda/refactory/client-server/api/master/repository"
)


type DataUsecaseImple struct {
	dataRepo repository.DataRepo
}

func InitDataUsecase(dataRepo repository.DataRepo) DataUsecase {
	return &DataUsecaseImple{dataRepo: dataRepo}
}

func (d DataUsecaseImple) PostData(data model.Data) error {
	err := d.dataRepo.CreateData(data)
	if err != nil {
		log.Println(err)
	}
	return err
}