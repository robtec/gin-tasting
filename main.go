package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robtec/gin-tasting/gins"
)

func main() {

	router := gin.Default()

	router.GET("/gins", listAllGins)

	router.Run(":8080")
}

func listAllGins(c *gin.Context) {

	gins := gins.All()

	c.JSON(http.StatusOK, gins)
}
