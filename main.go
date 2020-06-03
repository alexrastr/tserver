package main

import (
	"time"

	"github.com/cnjack/throttle"
	"github.com/gin-gonic/gin"
)

func main() {
	//перезапуск подписки при остутсвии данных в течении 3 минут
	mqttSubscribe(180000)

	router := gin.Default()

	//лимит запросов 5 в минуту
	router.Use(throttle.Policy(&throttle.Quota{
		Limit:  5,
		Within: time.Minute,
	}))

	router.GET("/api/v1/data", func(c *gin.Context) {
		sensor, err := getLastTemp()

		c.JSON(200, gin.H{
			"temp":  sensor.Temp,
			"date":  sensor.Model.CreatedAt,
			"error": err,
		})
	})
	router.Run("localhost:8080")
}
