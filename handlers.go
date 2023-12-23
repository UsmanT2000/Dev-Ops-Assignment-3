package main

import (
	"fmt"
	"github.com/UsmanT2000/ginAPIs/csv"
	"github.com/UsmanT2000/ginAPIs/models"
	"github.com/UsmanT2000/ginAPIs/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func handleParseCSV(ctx *gin.Context) {
	var fileRequest models.FileRequest

	err := utils.ReadJSONRequest(ctx, &fileRequest)
	if err != nil {
		utils.SendJSONResponse(ctx, nil, http.StatusBadRequest)
	}

	data, err := csv.ParseCSV(fileRequest.FilePath)
	if err != nil {
		utils.SendJSONResponse(ctx, data, http.StatusInternalServerError)
	}
	utils.SendJSONResponse(ctx, data, http.StatusOK)
}

func SayHello(ctx *gin.Context) {
	var request models.HelloRequest
	err := utils.ReadJSONRequest(ctx, &request)
	if err != nil {
		utils.SendJSONResponse(ctx, nil, http.StatusBadRequest)
	}
	// Process request and create response
	response := models.HelloResponse{
		Code:      200,
		Message:   fmt.Sprintf("Welcome %s!", request.Name),
		Timestamp: time.Now().UTC(),
	}

	utils.SendJSONResponse(ctx, response, http.StatusOK)
}
