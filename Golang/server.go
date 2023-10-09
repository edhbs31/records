package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func DocumentServer() {
	router := gin.New()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	err := router.Run(fmt.Sprintf("%s%s", ":", os.Getenv("PORT_DOCS")))
	if err != nil {
		log.Println("error starting server Documentations", err)
	}
}

func CatalogServer(router *gin.Engine) {
	s := &http.Server{
		Addr:         fmt.Sprintf("%s%s", ":", os.Getenv("PORT")),
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		// MaxHeaderBytes: 1 << 20,
		// ErrorLog:,
	}
	err := s.ListenAndServe()
	fmt.Print("this is errorr connect", err)
	if err != nil {

		// l.Error("Error starting server", "error", err)
		log.Println("error starting server ", err)
		os.Exit(1)
	}
	log.Println("Starting server Catalog  on port ", os.Getenv("PORT"))
}
