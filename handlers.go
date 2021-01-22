package main

import (
	"HushTell/model"
	"HushTell/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// UploadFileHandler handles file uploads
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers
	headers := w.Header()
	headers.Set("Content-Type", "application/json")
	headers.Set("Access-Control-Allow-Origin", "*")
	// Handle OPTIONS first
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// if r.Method == "POST"

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

	// get the file from request
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		log.Println(err)
		resp, _ := json.Marshal(map[string]string{"error": "request Content-Type isn't multipart/form-data"}) // no way there would be an error, so just ignore it
		w.Write(resp)
		return
	}
	defer file.Close()

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
	CachedFiles[tempFileName] = obj

	// response to the request
	resp, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

// UploadTextHandler handles notes uploads
func UploadTextHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers
	headers := w.Header()
	headers.Set("Content-Type", "application/json")
	headers.Set("Access-Control-Allow-Origin", "*")
	// Handle OPTIONS first
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// if r.Method == "POST"

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

	// get the text
	text := r.FormValue("mytext")

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
		Duration:   3 * time.Second}
	CachedFiles[tempFileName] = obj

	// response to the request
	resp, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

// FileAccessHandler handles file accesses
func FileAccessHandler(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS first
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// if r.Method == "GET"
	vars := mux.Vars(r)
	ext := CachedFiles[vars["hashedIP"]+"/"+vars["filename"]].Extension
	fmt.Fprintf(w, "The file you are trying to access is: %s\n", vars["filename"]+ext)
	path := "./temp/" + vars["hashedIP"] + "/" + vars["filename"] + ext
	folder := "./temp/" + vars["hashedIP"]
	_, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(w, "The file does not exist or has been destroyed.")
	} else {
		fmt.Fprintln(w, "Found the file.")
		aed := CachedFiles[vars["hashedIP"]+"/"+vars["filename"]].Duration
		fmt.Fprintln(w, "The file will destroy itself in ", aed)
		go util.InitAccessedTimer(vars["hashedIP"]+"/"+vars["filename"], folder, path, time.Now(), aed, &CachedFiles)
	}
}

// Index is just a page to let you know it's working
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>The server is running.</h1>")
}
