package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/sarahayunanda/refactory/client-server/api/master/model"
	"github.com/golang/sarahayunanda/refactory/client-server/api/master/usecase"
	"github.com/gorilla/mux"
)

type DataHandler struct {
	DataUsecase usecase.DataUsecase
}

func DataController(r *mux.Router, service usecase.DataUsecase) {
	dataHandler := DataHandler{service}
	r.HandleFunc("/blogs", dataHandler.PostingData).Methods(http.MethodPost)
}

func (d DataHandler) PostingData(w http.ResponseWriter, r *http.Request) {
	data := model.Data{}
	w.Header().Set("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

	err = d.DataUsecase.PostData(data, r)
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error"))
	} else {
		status := fmt.Sprintf("HTTP Status Code %d", http.StatusCreated)
		w.Write([]byte(status))
	}
}
