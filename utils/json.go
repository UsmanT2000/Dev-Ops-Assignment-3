package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// To Format output for the GET request
func SendJSONResponse(ctx *gin.Context, data interface{}, statusCode int) {
	ctx.Header("Content-Type", "application/json")
	ctx.Writer.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("JSON Encoding Failed")
	}

	ctx.Writer.Write(jsonData)
}

// Read JSON request body and decode into a struct
func ReadJSONRequest(c *gin.Context, data interface{}) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	defer c.Request.Body.Close()

	return json.Unmarshal(body, data)
}
