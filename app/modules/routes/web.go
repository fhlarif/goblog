package routes

import (
	"net/http"

	"github.com/fhlarif/goblog/app/pkg/renderer"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		renderer.View(c, http.StatusOK, "modules/view/index", gin.H{
			"title": "Index Page",
		})
	})

	r.GET("/about", func(c *gin.Context) {
		renderer.View(c, http.StatusOK, "modules/view/about", gin.H{
			"title": "About Page",
		})
	})
}
