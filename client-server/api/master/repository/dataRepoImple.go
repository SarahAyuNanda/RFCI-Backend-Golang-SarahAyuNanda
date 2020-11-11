package repository

import "github.com/golang/sarahayunanda/refactory/client-server/api/master/model"

type DataRepoImple struct {
}

func InitDataRepoImple() DataRepo {
	return &DataRepoImple{}
}

func (d *DataRepoImple) CreateData(data model.Data) error {
	return nil
}