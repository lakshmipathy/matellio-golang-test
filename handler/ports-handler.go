package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matellio-golang-test/model"
	"github.com/matellio-golang-test/util"
)

var err error

func GetAllPorts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(getMapDetails())
	util.HandleGenericError(err, "error in json encoding")
}
func GetPortByID(w http.ResponseWriter, r *http.Request) {
	requestParameters := mux.Vars(r)
	requestID := requestParameters["id"]
	w.Header().Add("Content-Type", "application/json")
	ports := getMapDetails()
	data, exist := ports[requestID]
	if exist {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(data)
		util.HandleGenericError(err, "error in json encoding")
	} else {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode("no record found for the id:   " + requestID)
		util.HandleGenericError(err, "error in json encoding")
	}
}
func CreateNewPort(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, ioerr := io.ReadAll(r.Body)
	requestParameters := mux.Vars(r)
	requestID := requestParameters["id"]
	if ioerr != nil {
		log.Fatal("error in reading requst ", ioerr)
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode("request json is not in valid format")
		util.HandleGenericError(err, "error in json encoding")
		return
	}
	var cityInfo model.CityInfo
	err = json.Unmarshal(body, &cityInfo)
	util.HandleGenericError(err, "error in unmarshaling")
	portData := getMapDetails()
	_, exist := portData[requestID]

	if exist {
		w.WriteHeader(http.StatusOK)
		portData[requestParameters["id"]] = cityInfo
		err = json.NewEncoder(w).Encode("record found for the id:  " + requestID + "updating the record")
		util.HandleGenericError(err, "error in json encoding")
	} else {
		w.WriteHeader(http.StatusCreated)
		portData[requestParameters["id"]] = cityInfo
		err = json.NewEncoder(w).Encode("new record created for the id:  " + requestID)
		util.HandleGenericError(err, "error in json encoding")
	}
}
func UpdatePort(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, ioerr := io.ReadAll(r.Body)
	requestParameters := mux.Vars(r)
	requestID := requestParameters["id"]
	if ioerr != nil {
		log.Fatal("error in reading requst ", ioerr)
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode("request json is not in valid format")
		util.HandleGenericError(err, "error in json encoding")
		return
	}
	var cityInfo model.CityInfo
	err = json.Unmarshal(body, &cityInfo)
	util.HandleGenericError(err, "error in json unmarshal")
	portData := getMapDetails()
	_, exist := portData[requestID]

	if exist {
		w.WriteHeader(http.StatusOK)
		portData[requestID] = cityInfo
		err = json.NewEncoder(w).Encode("record updated for id:  " + requestID)
		util.HandleGenericError(err, "error in json encoding")
	} else {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode("no record found for the id:  " + requestID)
		util.HandleGenericError(err, "error in json encoding")
	}
}
func DeletePortByID(w http.ResponseWriter, r *http.Request) {
	requestParameters := mux.Vars(r)
	requestID := requestParameters["id"]
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ports := getMapDetails()
	_, exist := ports[requestID]
	if exist {
		w.WriteHeader(http.StatusOK)
		delete(ports, requestID)
		err = json.NewEncoder(w).Encode("deleted the record for id:  " + requestID)
		util.HandleGenericError(err, "error in json encoding")
	} else {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode("no record found with the id:   " + requestID)
		util.HandleGenericError(err, "error in json encoding")
	}
}
