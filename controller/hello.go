package controller

import (
	"gin_demo/logger"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	logger.Logger.Info("hello world")
	c.JSON(200, gin.H{"code": 0, "msg": "hello world"})
	//c.String(http.StatusOK, "hello World!")
}
