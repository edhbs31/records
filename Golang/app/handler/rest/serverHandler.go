package rest

import (
	"log"
	"os"
	"record/app/config"
	"record/app/usecasecontainer"

	"github.com/gin-gonic/gin"
)

type serverHandler struct {
	useCase *usecasecontainer.UsecaseContainer
}

func NewServerHandler(r *gin.Engine, uc *usecasecontainer.UsecaseContainer) {
	handler := serverHandler{
		useCase: uc,
	}

	r.GET("/", handler.ServerCheck)
}

// @BasePath /server

// @Summary Check Server
// @Description Check Server
// @Tags CheckServer
// @Accept json
// @Produce json
// @Success 200 {object} domains.LoginResponse
// @Router / [get]
func (h serverHandler) ServerCheck(ctx *gin.Context) {
	scoop := config.Db{}
	buildConnection := scoop.Build()
	buildConnection.SetDsn(config.DsnScoop())
	conn := buildConnection.GetConnections()
	if conn.Error != nil {
		log.Print("error connecting to db!")
		os.Exit(1)
	}
}
