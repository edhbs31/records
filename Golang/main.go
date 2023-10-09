package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"record/app/config"
	"record/app/handler/rest"
	"record/app/repocontainer"
	"record/app/usecasecontainer"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// @version 1.0
// @description This is a sample for access api data access

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name api-key

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @schemes https http

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:  []string{"GET", "POST", "PUT"},
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "api-key", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	scoop := config.Db{}
	buildConnection := scoop.Build()
	buildConnection.SetDsn(config.DsnScoop())
	conn := buildConnection.GetConnections()
	//// set middleware basic access

	// set up container and usecase
	repo := repocontainer.NewRepoContainer(conn)
	usecase := usecasecontainer.NewUsecaseContainer(repo)

	//route server check
	rest.NewServerHandler(router, usecase)

	router.Use(gin.LoggerWithFormatter(LogFormatter))
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "ROUTER NOT FOUND"})
	})
	rG := router.Group("/")

	rest.NewDataHandler(rG, usecase)
	go CatalogServer(router)
	go DocumentServer() // server doscumentation
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

}
func LogFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage)
}
