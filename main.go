package main

import (
	"HushTell/config"
	"HushTell/model"
	"HushTell/util"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var tempFileTimers map[string]model.SavedFile = make(map[string]model.SavedFile)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS first
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// if r.Method == "POST"
	// Set headers
	headers := w.Header()
	headers.Set("Content-Type", "application/json")
	headers.Set("Access-Control-Allow-Origin", "*")

	// get the client IP
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	// hash the IP
	clientHash := util.ShortHash(clientIP)

	// create folder
	util.CreateFolderByName(clientHash)
	log.Println("Receiving an upload from: " + clientIP)

	// parse the form
	r.ParseMultipartForm(10 << 20)
	r.ParseForm()

	// get the file from request
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		log.Println(err)
		resp, _ := json.Marshal(map[string]string{"error": "request Content-Type isn't multipart/form-data"})
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
	obj := model.SavedFile{
		Filename:               handler.Filename,
		Url:                    "f/" + tempFileName,
		Extension:              ext,
		InitTime:               time.Now(),
		ExpireDuration:         config.GlobalExpireDuration,
		AccessedExpireDuration: 3 * time.Second}
	tempFileTimers[tempFileName] = obj

	// response to the request
	resp, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

func uploadTextHandler(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS first
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// if r.Method == "POST"
	// Set headers
	headers := w.Header()
	headers.Set("Content-Type", "application/json")
	headers.Set("Access-Control-Allow-Origin", "*")

	// get the client IP
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	// hash the IP
	clientHash := util.ShortHash(clientIP)

	// create folder
	util.CreateFolderByName(clientHash)
	log.Println("Receiving an upload from: " + clientIP)

	// parse the form
	r.ParseMultipartForm(10 << 20)
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
	obj := model.SavedFile{
		Filename:               "",
		Url:                    "f/" + tempFileName,
		Extension:              ".txt",
		InitTime:               time.Now(),
		ExpireDuration:         config.GlobalExpireDuration,
		AccessedExpireDuration: 3 * time.Second}
	tempFileTimers[tempFileName] = obj

	// response to the request
	resp, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Error parsing template.")
		log.Println(err)
		return
	}
	data := map[string]string{"PORT": config.PORT}
	tmpl.Execute(w, data)
}

func fileAccessHandler(w http.ResponseWriter, r *http.Request) {
	// Handle OPTIONS first
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// if r.Method == "GET"
	vars := mux.Vars(r)
	ext := tempFileTimers[vars["hashedIP"]+"/"+vars["filename"]].Extension
	fmt.Fprintf(w, "The file you are trying to access is: %s\n", vars["filename"]+ext)
	path := "./temp/" + vars["hashedIP"] + "/" + vars["filename"] + ext
	folder := "./temp/" + vars["hashedIP"]
	_, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(w, "The file does not exist or has been destroyed.")
	} else {
		fmt.Fprintln(w, "Found the file.")
		aed := tempFileTimers[vars["hashedIP"]+"/"+vars["filename"]].AccessedExpireDuration
		fmt.Fprintln(w, "The file will destroy itself in ", aed)
		go util.InitAccessedTimer(vars["hashedIP"]+"/"+vars["filename"], folder, path, time.Now(), aed, &tempFileTimers)
	}
}

func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/upload-file", uploadFileHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/upload-text", uploadTextHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/f/{hashedIP}/{filename}", fileAccessHandler).Methods("GET", "OPTIONS")
	http.ListenAndServe(":"+config.PORT, r)
}

func checkGlobalTimer() {
	for 1 == 1 {
		for key := range tempFileTimers {
			if time.Now().Sub(tempFileTimers[key].InitTime) > tempFileTimers[key].ExpireDuration {
				os.Remove(tempFileTimers[key].Filename)
			}
		}
		time.Sleep(3 * time.Second)
		log.Println(tempFileTimers)
	}
}

func main() {
	log.Printf("Running a simple server at port %s...\n", config.PORT)
	go checkGlobalTimer()
	setupRoutes()
}
