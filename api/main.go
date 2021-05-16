package musiclink

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rapito/go-spotify/spotify"
)

func main() {
	r := gin.Default()
	fmt.Println("running api...")

	err := godotenv.Load("prod.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	spot := spotify.New(clientID,clientSecret)
    result, _ := spot.Get("albums/%s", nil, "0sNOF9WDwhWunNAHPD3Baj")
	fmt.Println(result)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.Run()
}