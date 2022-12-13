package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HttpResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, ApiResult{code, message, data})
}
