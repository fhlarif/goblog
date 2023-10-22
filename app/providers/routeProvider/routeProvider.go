package routeprovider

import (
	web "github.com/fhlarif/goblog/app/resources/routes"
	"github.com/gin-gonic/gin"
)

func RoutesProvider(router *gin.Engine) {
	web.Routes(router)
}
