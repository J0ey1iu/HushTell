// Package handler API documentation
//
// Documentation for tests
//
// Schemes: http
// Host: localhost:8000
// BasePath: /
// Version: 0.0.1
//
// Consumes:
// - multipart/form-data
//
// Produces:
// - application/json
//
// swagger:meta
package handler

import (
	"HushTell/model"
	"HushTell/util"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// FileAccessHandler handles file accesses
func FileAccessHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /f/{hashedIP}/{filename} access getFile
	//
	// locate and try to access file and return the result,
	// if the file is found, the timer will start run and automatically destroy that file
	//
	// Consumes:
	// - application/json
	//
	// Produces:
	// - application/json
	//
	// Schemes: http
	//
	// Responses:
	// - 200: AccessResult

	vars := mux.Vars(r)
	ext := model.CachedFiles[vars["hashedIP"]+"/"+vars["filename"]].Extension
	fmt.Fprintf(w, "The file you are trying to access is: %s\n", vars["filename"]+ext)
	path := "./temp/" + vars["hashedIP"] + "/" + vars["filename"] + ext
	folder := "./temp/" + vars["hashedIP"]
	_, err := os.Stat(path)
	if err != nil {
		resp, _ := json.Marshal(map[string]string{
			"result": "The file does not exist or has been destroyed.",
		})
		w.Write(resp)
	} else {
		aed := model.CachedFiles[vars["hashedIP"]+"/"+vars["filename"]].Duration
		resp, _ := json.Marshal(map[string]string{
			"result": "Found the file, timer is running...",
		})
		w.Write(resp)
		go util.InitAccessedTimer(vars["hashedIP"]+"/"+vars["filename"], folder, path, time.Now(), aed, &model.CachedFiles)
	}
}
