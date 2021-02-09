package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//in go functions and structs that start lowercase cannot be acccessed outside that folder

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{})
}
