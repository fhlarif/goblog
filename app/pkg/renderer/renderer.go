package renderer

import (
	viewprovider "github.com/fhlarif/goblog/app/providers/viewProvider"
	"github.com/gin-gonic/gin"
)

func LoadHTML(router *gin.Engine) {
	router.LoadHTMLGlob("app/**/**/*.tmpl")
}

func LoadStatic(router *gin.Engine) {
	router.Static("/app/assets", "./app/assets")
}

func View(c *gin.Context, code int, page string, data gin.H) {
	data = viewprovider.ViewProvider(data)
	c.HTML(code, page, data)
}
