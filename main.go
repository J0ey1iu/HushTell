package main

import (
	"HushTell/config"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Page after upload</h1>\n")
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data.")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("Uploaded size: %+v\n", handler.Size)

	ext := filepath.Ext(handler.Filename)
	tempFile, err := ioutil.TempFile("temp", "upload-*"+ext)
	if err != nil {
		fmt.Println("Error creating temp file.")
		fmt.Println(err)
		return
	}
	defer tempFile.Close()
	tempFileName := tempFile.Name()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading uploaded file.")
		fmt.Println(err)
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		fmt.Println("Error writing the temp file.")
		fmt.Println(err)
		return
	}

	fmt.Fprintln(w, "Uploaded successfully.")
	fmt.Fprintf(w, "Corresponding file <a href=\"%s\">link", "http://localhost:"+config.PORT+"/"+tempFileName)
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
	_, err := os.Stat("./temp/" + vars["filename"])
	if err != nil {
		fmt.Fprintln(w, "The file does not exist.")
	} else {
		fmt.Fprintln(w, "Found the file.")
	}
}

func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadHandler)
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/temp/{filename}", fileAccessHandler)
	http.ListenAndServe(":"+config.PORT, r)
}

func main() {
	fmt.Printf("Running a simple server at port %s...", config.PORT)
	setupRoutes()
}
