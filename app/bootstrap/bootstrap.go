package bootstrap

import (
	"github.com/fhlarif/goblog/app/pkg/config"
	"github.com/fhlarif/goblog/app/pkg/renderer"
	"github.com/fhlarif/goblog/app/pkg/router"
)

func InitBootstrap() {
	config.SetConfig()
	router.Init()
	renderer.LoadStatic(router.GetRouter())
	renderer.LoadHTML(router.GetRouter())
	router.RegisterRouter()
	router.SetRouter()
}
