package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"net/http"
)

type Task struct {
	ID     int
	Name   string
	Status bool
}

const (
	Port = "8080"
)

/*func init() {
	//db.Connect()
}*/

func main() {
	router := gin.Default()

	//router.Use(middlewares.ErrorHandler)
	//router.Static("/public", "./public")

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	http.Handle("/", router)

	//router.Run(":" + Port)
	appengine.Main()
}
