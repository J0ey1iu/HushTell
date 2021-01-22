package main

import (
	"HushTell/config"
	"HushTell/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// CachedFiles A cache storing all the files uploaded
var CachedFiles map[string]model.CachedInfo = make(map[string]model.CachedInfo)

func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/upload-file", UploadFileHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/upload-text", UploadTextHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/f/{hashedIP}/{filename}", FileAccessHandler).Methods("GET", "OPTIONS")
	http.ListenAndServe(":"+config.PORT, r)
}

func checkGlobalTimer() {
	for 1 == 1 {
		for key := range CachedFiles {
			if time.Now().Sub(CachedFiles[key].UploadTime) > config.GlobalExpireDuration {
				os.Remove(CachedFiles[key].Filename)
			}
		}
		time.Sleep(config.GlobalCheckRate)
		log.Println(CachedFiles)
	}
}

func main() {
	log.Printf("Running a simple server at port %s...\n", config.PORT)
	go checkGlobalTimer()
	setupRoutes()
}
