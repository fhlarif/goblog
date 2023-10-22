package router

import (
	"fmt"
	"log"

	"github.com/fhlarif/goblog/app/pkg/config"
	routeprovider "github.com/fhlarif/goblog/app/providers/routeProvider"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func SetRouter() {
	r := GetRouter()
	configs := config.GetConfig()
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}

func RegisterRouter() {
	routeprovider.RoutesProvider(GetRouter())
}
