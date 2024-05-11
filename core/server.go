package core

import (
	"dsvm/global"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

const StaticDir = "./static"

func RunHttpServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery(), cors.Default())
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFileFS("/favicon.ico", "favicon.ico", http.Dir("static"))
	router.StaticFS("/vs", http.Dir("static/monaco-editor/vs"))
	router.Delims("[{[", "]}]")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		dir, _ := os.ReadDir(MigrationsDir)
		var fileNames = make(map[string]string, len(dir))
		for _, entry := range dir {
			key := fmt.Sprintf("%s/%s", MigrationsDir, entry.Name())
			fileNames[key[1:]] = entry.Name()
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":     "Database Schema Version Management",
			"fileNames": fileNames,
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}
		files := form.File["files"]
		index := ReadMigrationsIndex()
		for _, file := range files {
			index = index + 1
			dst := fmt.Sprintf("%s/%d_%s", MigrationsDir, index, file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}
		WriteMigrationsIndex(index)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	router.Any("/up", func(c *gin.Context) {
		params := struct {
			Schema string `json:"schema" form:"schema"`
			DS     string `json:"ds" form:"ds"`
		}{}
		_ = c.ShouldBind(&params)
		err := UpBySchema(params.DS, params.Schema)
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
