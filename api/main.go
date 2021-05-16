package main

import (
	"fmt"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"main/controllers"
	//"reflect"
)

func main() {
	r := gin.Default()
	fmt.Println("running api...")

	err := godotenv.Load("prod.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))


	sptfyWrapper := controllers.NewSpotifyController()
	r.GET("/spotify/tracks/:id", sptfyWrapper.GetTrack)
	r.Run()
}