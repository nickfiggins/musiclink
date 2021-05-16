package controllers

import (
	"main/sptfy"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SpotifyController struct {
	SpotifyAPI *sptfy.SpotifyAPI
}

func NewSpotifyController() *SpotifyController{
	apiWrapper := sptfy.NewSpotifyAPI()
	return &SpotifyController{SpotifyAPI: apiWrapper}
}


// /spotify/track/:id
func (sCtrl *SpotifyController) GetTrack(c *gin.Context) {
	track, err := sCtrl.SpotifyAPI.GetTrack(c.Param("id")); if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		c.AbortWithStatus(http.StatusForbidden)
	}

	c.JSON(http.StatusAccepted, track)
}