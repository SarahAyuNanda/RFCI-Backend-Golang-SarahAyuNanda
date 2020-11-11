package repository

import (
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
	aut := fmt.Sprintf("%v Success POST http://%v%v {\"author\" : \"%v\"} \n", time.Now(), r.Host, r.URL.Path, data.Author )
	if _, err := author.WriteString(aut); err != nil {
		log.Println(err)
	}

	title, err := os.OpenFile("../../log/title.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	ttl := fmt.Sprintf("%v Success POST http://%v%v {\"title\" : \"%v\"} \n", time.Now(), r.Host, r.URL.Path, data.Title )
	if _, err := title.WriteString(ttl); err != nil {
		log.Println(err)
	}

	message, err := os.OpenFile("../../log/message.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	msg := fmt.Sprintf("%v Success POST http://%v%v {\"comments\" : \"%v\"} \n", time.Now(), r.Host, r.URL.Path, data.Comment )
	if _, err := message.WriteString(msg); err != nil {
		log.Println(err)
	}

	return nil
}
