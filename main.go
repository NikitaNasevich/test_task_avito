package main

import (
	"github.com/NikitaNasevich/test_task_avito/log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err = r.Run(":8080"); err != nil {
		log.GetLogger().Errorf("Can't start server: %s", err.Error())
		os.Exit(1)
	}
}
