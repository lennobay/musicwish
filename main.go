package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"musicwish.haufe-server.de/api"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "wish"
	password = "wish"
	dbname   = "wishdb"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api.CreateNewEvent()
	api.CreateSessionDB()
	router.GET("/api/status", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.POST("/api/admin/login", api.Login)
	router.GET("/api/admin/delete_song", api.RemoveSong)
	router.GET("/api/client/all_songs", api.AllSongsOfEvent)
	router.POST("/api/client/add_song", api.AddNewSong)
	router.POST("/api/client/vote_song", api.VoteforSong)
	router.NoRoute(func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.Dir("./website"))
	})
	router.Run("localhost:8081")
}
