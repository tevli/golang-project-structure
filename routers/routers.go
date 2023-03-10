package routers

import (
	"github.com/gin-gonic/gin"
	apiControllerV1 "mklogistics/controllers/api/v1"
	apiControllerV2 "mklogistics/controllers/api/v2"
	webhandler "mklogistics/controllers/web"
	"mklogistics/middlewares"
)

//SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {

	r := gin.Default()

	//Giving access to storage folder
	r.Static("/storage", "storage")

	//Giving access to template folder
	r.Static("/templates", "templates")
	r.LoadHTMLGlob("templates/*")

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//API route for version 1
	v1 := r.Group("/api/v1")

	//Web route for web views.
	web := r.Group("web")
	web.GET("demo.html", webhandler.ViewDemo)

	//If you want to pass your route through specific middlewares
	v1.Use(middlewares.UserMiddlewares())
	{
		v1.POST("user-list", apiControllerV1.UserList)
	}

	//API route for version 2
	v2 := r.Group("/api/v2")

	v2.POST("user-list", apiControllerV2.UserList)

	return r

}
