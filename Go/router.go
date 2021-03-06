package main

import (
	"github.com/gin-gonic/gin"
	. "raspberry/apis"
	"raspberry/middleware"
	"raspberry/models"
)

func initRouter(u *models.UPS) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.GET("/", IndexAPI)
	router.GET("/battery", u.BatteryAPI)
	router.GET("/wallpaper/:resolution", WallpaperIndex)
	return router
}
