package calendar

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

//in go functions and structs that start lowercase cannot be acccessed outside that folder

func Event(c *gin.Context) {

	event := c.Param("id")

	fmt.Println(event)

	c.HTML(http.StatusOK, "calendar/select-day-and-time.html", gin.H{})
}
