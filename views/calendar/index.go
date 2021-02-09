package calendar

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//in go functions and structs that start lowercase cannot be acccessed outside that folder

func Index(c *gin.Context) {
	user := c.Param("user")

	c.HTML(http.StatusOK, "calendar/index.html", gin.H{"User": user})
}
