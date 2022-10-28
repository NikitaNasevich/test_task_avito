package main

import (
	"github.com/NikitaNasevich/test_task_avito/api"
	"github.com/NikitaNasevich/test_task_avito/db"
	"github.com/NikitaNasevich/test_task_avito/helpers"
	"github.com/NikitaNasevich/test_task_avito/log"
	"github.com/joho/godotenv"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	_ = godotenv.Load()
	var err error

	if err = db.ConnectDatabase(); err != nil {
		log.GetLogger().Errorf("Can't connect to database: %s", err.Error())
		os.Exit(1)
	}

	r := gin.Default()

	api.AllApi(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := helpers.GetEnvDefault("PORT", "3000")

	if err = r.Run(":" + port); err != nil {
		log.GetLogger().Errorf("Can't start server: %s", err.Error())
		os.Exit(1)
	}
}
