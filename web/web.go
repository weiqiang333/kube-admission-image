package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/weiqiang333/kube-admission-image/web/handlers"
)

func Web() {
	address := viper.GetString("address")
	tls := viper.GetBool("tls")
	cert := viper.GetString("cert")
	key := viper.GetString("key")

	router := gin.Default()

	router.GET("/healthz", func(context *gin.Context) {
		context.String(http.StatusOK, "healthz is OK")
	})
	router.POST("/images_admission", handlers.ImagesAdmission)

	if tls {
		router.RunTLS(address, cert, key)
	}
	router.Run(address)
}
