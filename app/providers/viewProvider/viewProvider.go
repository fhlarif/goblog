package viewprovider

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ViewProvider(data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")

	return data
}
