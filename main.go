package main

import (
	"HushTell/config"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the upload page.")
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
	tempFile, err := ioutil.TempFile("temp", "upload-*."+ext)
	if err != nil {
		fmt.Println("Error creating temp file.")
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

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

func setupRoutes() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", indexPage)
	http.ListenAndServe(":"+config.PORT, nil)
}

func main() {
	fmt.Printf("Running a simple server at port %s...", config.PORT)
	setupRoutes()
}
