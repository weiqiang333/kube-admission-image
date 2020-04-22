package web

import (
	"net/http"

	"github.com/weiqiang333/kube-admission-image/web/handlers"

	"github.com/gin-gonic/gin"

	"github.com/weiqiang333/kube-admission-image/pkg/config"
)

func Web(configs config.FlagVar) {
	router := gin.Default()

	router.GET("/healthz", func(context *gin.Context) {
		context.String(http.StatusOK, "healthz is OK")
	})
	router.POST("/images_admission", handlers.ImagesAdmission)

	if configs.Tls {
		router.RunTLS(configs.Addrss, configs.CertFile, configs.KeyFile)
	}
	router.Run(configs.Addrss)
}
