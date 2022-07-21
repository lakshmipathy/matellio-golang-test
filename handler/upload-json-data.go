package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/matellio-golang-test/config"
	"github.com/matellio-golang-test/model"
	"github.com/matellio-golang-test/util"
)

var result map[string]model.CityInfo

func UploadJSONData() {
	filename := config.GetUploadFilePath()
	f, ioerr := ioutil.ReadFile(filename)
	if ioerr != nil {
		log.Fatal("error in reading file: ", ioerr)
		return
	}
	err = json.Unmarshal([]byte(f), &result)
	util.HandleGenericError(err, "error in json unmarshal")
}

func getMapDetails() map[string]model.CityInfo {
	return result
}
