package rest

import (
	"fmt"
	"net/http"
	"record/app/domains"
	"record/app/usecasecontainer"

	"github.com/gin-gonic/gin"
)

type DataHandler struct {
	usecase *usecasecontainer.UsecaseContainer
}

func NewDataHandler(r *gin.RouterGroup, uc *usecasecontainer.UsecaseContainer) {
	handler := &DataHandler{
		usecase: uc,
	}
	data := r.Group("/data") // proxy nginx data
	data.GET("/", handler.FindAll)
	data.POST("/", handler.InsertData)
}

// @BasePath /data

// @Summary data example
// @Schemes
// @Description get data by id
// @Tags data
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {array} domains.StandartHttpResponse
// @Router /data?startDate=&endDate=&minCount=&maxCount= [get]
func (l DataHandler) FindAll(ctx *gin.Context) {
	var payload domains.DataRequest
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	data, err := l.usecase.DataUsecase.FindAll(payload.StartDate, payload.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "data not found",
			"error":  err.Error(),
		})
		return
	}
	response, err := l.usecase.DataUsecase.FilterData(data, payload.MaxCount, payload.MinCount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed filter data",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response,
	})
}

// @BasePath /data

// @Summary data example
// @Schemes
// @Description get data by id
// @Tags data
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {array} domains.StandartHttpResponse
// @Router /data {"name": "", "marks": [ ] }
func (l DataHandler) InsertData(ctx *gin.Context) {
	var payload domains.CasRecord
	if err := ctx.ShouldBind(&payload); err != nil {
		fmt.Print("pay load error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	fmt.Print("payload", payload)
	data, err := l.usecase.DataUsecase.InsertData(payload)
	fmt.Print("this is res", data, err)
}
