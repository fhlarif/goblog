package routeprovider

import (
	web "github.com/fhlarif/goblog/app/modules/routes"
	"github.com/gin-gonic/gin"
)

func RoutesProvider(router *gin.Engine) {
	web.Routes(router)
}
