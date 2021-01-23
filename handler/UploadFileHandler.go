package handler

import (
	"HushTell/model"
	"HushTell/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// UploadFileHandler handles file uploads
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /api/upload-file upload uploadFile
	//
	// take the user input and upload the file
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
	r.ParseMultipartForm(10 << 20) // 10MB
	r.ParseForm()
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		log.Println(err)
		resp, _ := json.Marshal(map[string]string{"error": "request Content-Type isn't multipart/form-data"}) // no way there would be an error, so just ignore it
		w.Write(resp)
		return
	}
	defer file.Close()
	optionsString := r.FormValue("options")
	var options model.Option
	json.Unmarshal([]byte(optionsString), &options)

	// create a tempfile and write the info down
	ext := filepath.Ext(handler.Filename)
	tempFile, err := ioutil.TempFile("temp/"+clientHash, "*"+ext)
	defer tempFile.Close()
	log.Println("Creating tempFile at: " + tempFile.Name())
	tempFileName := strings.Split(strings.Join(strings.Split(tempFile.Name(), "/")[1:], "/"), ".")[0]
	fileBytes, err := ioutil.ReadAll(file)
	_, err = tempFile.Write(fileBytes)

	// create a new record
	obj := model.CachedInfo{
		Filename:   handler.Filename,
		Url:        "f/" + tempFileName,
		Extension:  ext,
		UploadTime: time.Now(),
		Duration:   3 * time.Second}
	model.CachedFiles[tempFileName] = obj

	// response to the request
	resp, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}
