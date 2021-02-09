package main

import (
	"path/filepath"
	"strings"

	"net/http"

	"github.com/endingwithali/gogocal/configuration"
	"github.com/endingwithali/gogocal/views/calendar"
	"github.com/endingwithali/gogocal/views/home"

	"github.com/gin-contrib/multitemplate"
	"github.com/yargevad/filepathx"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Setup a new Zap logger.
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Starting GoGoCalendar!")

	// Setup our Gin instance.
	router := gin.Default()

	// Enable session storage
	store := cookie.NewStore([]byte(configuration.SessionKey))

	store.Options(sessions.Options{
		Domain: configuration.Host,
		Path:   "/",
	})

	// Configure session storage
	router.Use(sessions.Sessions("session", store))

	// Load our HTML templates
	router.HTMLRender = loadTemplates("templates")

	// Load our static assets.
	router.Static("/assets", "dist")

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state": "Success",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	home.Init(router)
	calendar.Init(router)

	err := router.Run(":5000")

	if err != nil {
		panic(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepathx.Glob(templatesDir + "/views/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and views/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(strings.ReplaceAll(include, "templates/views/", ""), files...)
	}

	return r
}
