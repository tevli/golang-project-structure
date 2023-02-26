package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ViewDemo(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.html", "")
}
func ViewIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}
func ViewAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", "")
}
