package main

import (
	"HushTell/config"
	"HushTell/handler"
	"HushTell/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"
	ghandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CachedFiles A cache storing all the files uploaded
// var CachedFiles map[string]model.CachedInfo = make(map[string]model.CachedInfo)

func setupRoutes() {
	r := mux.NewRouter()

	// GET router
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", handler.Index)
	getRouter.HandleFunc("/f/{hashedIP}/{filename}", handler.FileAccessHandler)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	h := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", h)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// POST router
	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/upload-file", handler.UploadFileHandler)
	postRouter.HandleFunc("/api/upload-text", handler.UploadTextHandler)

	// CORS
	ch := ghandler.CORS(ghandler.AllowedOrigins([]string{"*"})) // CORS handler

	http.ListenAndServe(":"+config.PORT, ch(r))
}

func checkGlobalTimer() {
	for 1 == 1 {
		for key := range model.CachedFiles {
			if time.Now().Sub(model.CachedFiles[key].UploadTime) > config.GlobalExpireDuration {
				os.Remove(model.CachedFiles[key].Filename)
			}
		}
		time.Sleep(config.GlobalCheckRate)
		log.Println(model.CachedFiles)
	}
}

func main() {
	log.Printf("Running a simple server at port %s...\n", config.PORT)
	go checkGlobalTimer()
	setupRoutes()
}
