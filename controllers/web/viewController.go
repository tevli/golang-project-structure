package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewDemo(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", "")
}
