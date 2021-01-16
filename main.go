package main

import (
	"HushTell/config"
	"HushTell/model"
	"HushTell/util"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var tempFileTimers map[string]model.SavedFile = make(map[string]model.SavedFile)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// get the client IP
	clientIP := strings.Split(r.RemoteAddr, ":")[0]
	// hash the IP
	clientHash := util.ShortHash(clientIP)
	// create folder
	util.CreateFolderByName(clientHash)

	// display on page
	fmt.Println("Receiving an upload from: " + clientIP)
	fmt.Fprintf(w, "<h1>Page after upload</h1>\n")
	r.ParseMultipartForm(10 << 20)

	// get the file from request
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// create a tempfile and write the info down
	ext := filepath.Ext(handler.Filename)
	tempFile, err := ioutil.TempFile("temp/"+clientHash, "*"+ext)
	defer tempFile.Close()
	fmt.Println("Creating tempFile at: " + tempFile.Name())
	tempFileName := strings.Join(strings.Split(tempFile.Name(), "/")[1:], "/")
	fileBytes, err := ioutil.ReadAll(file)
	_, err = tempFile.Write(fileBytes)

	// create a new record
	tempFileTimers[tempFileName] = model.SavedFile{Filename: tempFileName, InitTime: time.Now(), ExpireDuration: config.GlobalExpireDuration, AccessedExpireDuration: 3 * time.Second}

	fmt.Fprintln(w, "Uploaded successfully.")
	fmt.Fprintf(w, "Corresponding file <a href=\"%s\">link", "http://localhost:"+config.PORT+"/f/"+tempFileName)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Error parsing template.")
		fmt.Println(err)
		return
	}
	data := map[string]string{"PORT": config.PORT}
	tmpl.Execute(w, data)
}

func fileAccessHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "The file you are trying to access is: %s\n", vars["filename"])
	path := "./temp/" + vars["hashedIP"] + "/" + vars["filename"]
	folder := "./temp/" + vars["hashedIP"]
	_, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(w, "The file does not exist.")
	} else {
		fmt.Fprintln(w, "Found the file.")
		aed := tempFileTimers[vars["hashedIP"]+"/"+vars["filename"]].AccessedExpireDuration
		fmt.Fprintln(w, "The file will destroy itself in ", aed)
		go util.InitAccessedTimer(vars["hashedIP"]+"/"+vars["filename"], folder, path, time.Now(), aed, &tempFileTimers)
	}
}

func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadHandler)
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/f/{hashedIP}/{filename}", fileAccessHandler)
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
		fmt.Println(tempFileTimers)
	}
}

func main() {
	fmt.Printf("Running a simple server at port %s...\n", config.PORT)
	go checkGlobalTimer()
	setupRoutes()
}
