package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/sarahayunanda/refactory/client-server/api/master/model"
)

type DataRepoImple struct {
}

func InitDataRepoImple() DataRepo {
	return &DataRepoImple{}
}

func (d *DataRepoImple) CreateData(data model.Data, r *http.Request) error {
	author, err := os.OpenFile("../../log/author.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	objAut := model.Author{data.Author}
	jsonAut, _ := json.Marshal(objAut)
	jsonStringAut := string(jsonAut)
	aut := fmt.Sprintf("%v Success POST http://%v%v %s \n", time.Now(), r.Host, r.URL.Path, jsonStringAut)
	if _, err := author.WriteString(aut); err != nil {
		log.Println(err)
	}

	title, err := os.OpenFile("../../log/title.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	objTtl := model.Title{data.Title}
	jsonTtl, _ := json.Marshal(objTtl)
	jsonStringTtl := string(jsonTtl)
	ttl := fmt.Sprintf("%v Success POST http://%v%v %s \n", time.Now(), r.Host, r.URL.Path, jsonStringTtl)
	if _, err := title.WriteString(ttl); err != nil {
		log.Println(err)
	}

	message, err := os.OpenFile("../../log/message.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	objMsg := model.Comment{data.Comment}
	jsonMsg, _ := json.Marshal(objMsg)
	jsonStringMsg := string(jsonMsg)
	msg := fmt.Sprintf("%v Success POST http://%v%v %s \n", time.Now(), r.Host, r.URL.Path, jsonStringMsg)
	if _, err := message.WriteString(msg); err != nil {
		log.Println(err)
	}

	return nil
}
