package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"isithalloween.com/server"
)

var err error

func main() {
	err = godotenv.Load(".env")
	PORT := "8080"
	if err != nil {
		log.Println(".env file does not exist")
	}
	if envPORT := os.Getenv("PORT"); len(envPORT) > 0 {
		PORT = envPORT
	}
	router := gin.Default()

	router.Static("/", "./client")

	go receive(router)

	router.Run("localhost:" + PORT)
}

func receive(router *gin.Engine) {
	router.POST("/api", halloweenCheck)
}

func halloweenCheck(ctx *gin.Context) {
	reqbody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal("Request body reeading error:", err)
	}
	reqdate := &server.Date{Str: ""}
	if err := json.Unmarshal(reqbody, &reqdate); err != nil {
		ctx.JSON(http.StatusBadRequest, server.Message{"Error 400, Bad Request"})
	}
	res := "No"
	if server.IsHalloween(reqdate) {
		res = "Yes"
	}
	ctx.JSON(http.StatusOK, server.Message{res})
}
