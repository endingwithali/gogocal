package calendar

import "github.com/gin-gonic/gin"

func Init(router *gin.Engine) {
	routes := router.Group("/calendar")

	routes.GET("/:user", Index)
	routes.GET("/:user/:event", Event)
}
