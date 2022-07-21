/*
 */
package main

import (
	"github.com/matellio-golang-test/config"
	"github.com/matellio-golang-test/handler"
	httprouter "github.com/matellio-golang-test/http-router"
	"github.com/matellio-golang-test/util"
)

var err error

func main() {
	err = config.InitConfig("matellio", nil)
	util.HandleGenericError(err, "error in config init")
	handler.UploadJSONData()
	err = httprouter.StartHTTPServer()
	util.HandleGenericError(err, "error in starting http server")
}
