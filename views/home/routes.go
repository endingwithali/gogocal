package home

import "github.com/gin-gonic/gin"

func Init(router *gin.Engine) {
	routes := router.Group("/")

	routes.GET("/", Index)

}
