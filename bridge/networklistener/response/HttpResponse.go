package httpResponse

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateSuccessfulResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, ApiResponse{Status: code, Message: message, Data: data})
}
