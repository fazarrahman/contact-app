package main

import (
	"log"
	"net/http"

	sqllitedb "github.com/fazarrahman/contact-app/config"
	contactSqllite "github.com/fazarrahman/contact-app/contact/repository/sqllite"
	"github.com/fazarrahman/contact-app/controller"
	"github.com/fazarrahman/contact-app/lib"
	"github.com/fazarrahman/contact-app/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sqllitedb.New()
	if err != nil {
		panic(err)
	}

	contactRepo := contactSqllite.New(db)
	svc := service.New(contactRepo)

	//auth.Init()
	log.Println("Oauth has been initialized")

	g := gin.Default()
	g.GET("/ping",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	controller.New(svc).Register(g.Group("/api/v1"))

	g.Run(":" + lib.GetEnv("APP_PORT"))
}
