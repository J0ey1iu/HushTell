package handler

import (
	"HushTell/model"
	"HushTell/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// UploadTextHandler handles notes uploads
func UploadTextHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /api/upload-text upload uploadText
	//
	// upload the text user typed in along with the options they set
	//
	// Schemes: http
	//
	// Consumes:
	// - multipart/form-data
	//
	// Produces:
	// - application/json
	//
	// Responses:
	// - 200: CachedInfo

	// get the client IP
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	// hash the IP
	clientHash := util.ShortHash(clientIP)

	// create folder
	util.CreateFolderByName(clientHash)
	log.Println("Receiving an upload from: " + clientIP)

	// parse the form
	// TODO: we don't need multipart here
	r.ParseMultipartForm(10 << 20) // 10MB
	r.ParseForm()
	text := r.FormValue("mytext")
	optionsString := r.FormValue("options")
	var options model.Option
	json.Unmarshal([]byte(optionsString), &options)

	// create a tempfile and write the info down
	tempFile, err := ioutil.TempFile("temp/"+clientHash, "*.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer tempFile.Close()
	log.Println("Creating tempFile at: " + tempFile.Name())
	tempFileName := strings.Split(strings.Join(strings.Split(tempFile.Name(), "/")[1:], "/"), ".")[0]
	_, err = tempFile.Write([]byte(text))

	// create a new record
	obj := model.CachedInfo{
		Filename:   "",
		Url:        "f/" + tempFileName,
		Extension:  ".txt",
		UploadTime: time.Now(),
		Duration:   3 * time.Second,
		Options:    options}
	model.CachedFiles[tempFileName] = obj

	// response to the request
	resp, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}
