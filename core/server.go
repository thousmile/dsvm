package core

import (
	"dsvm/global"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func RunHttpServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery(), cors.Default())
	router.Any("/up", func(c *gin.Context) {
		params := struct {
			Schema  string `json:"schema" form:"schema"`
			Version int64  `json:"version" form:"version"`
			DS      string `json:"ds" form:"ds"`
		}{}
		_ = c.ShouldBind(&params)
		if params.Version == 0 {
			params.Version = 1
		}
		err := UpBySchema(params.DS, params.Schema, params.Version)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "ok",
			})
		}
	})

	router.Any("/down", func(c *gin.Context) {
		params := struct {
			Version int64  `json:"version" form:"version"`
			DS      string `json:"ds" form:"ds"`
		}{}
		_ = c.ShouldBind(&params)
		if params.Version == 0 {
			params.Version = 1
		}
		err := DownBySchema(params.DS, params.Version)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "ok",
			})
		}
	})

	router.Any("/version", func(c *gin.Context) {
		v1 := fmt.Sprintf("%s -> %s", global.AppConfig.AppName, global.AppConfig.Version)
		c.JSON(http.StatusOK, gin.H{
			"version": v1,
		})
	})
	router.Any("/health", func(c *gin.Context) {
		v7, _ := uuid.NewV7()
		c.JSON(http.StatusOK, gin.H{
			"message": v7.String(),
		})
	})
	log.Printf("http server listen : %d \n", global.AppConfig.HttpPort)
	addr := fmt.Sprintf(":%d", global.AppConfig.HttpPort)
	if err := router.Run(addr); err != nil {
		log.Panicf("http server listen error: %v", err)
	}
}
