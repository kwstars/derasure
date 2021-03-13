package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kwstars/derasure/internal/app/services"
	transporthttp "github.com/kwstars/derasure/pkg/transports/http"
	"net/http"
)

type Controller struct {
	Service services.Services
}

func CreateInitControllersFn(f *Controller) transporthttp.InitControllers {
	return func(r *gin.Engine) {

		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{})
		})

		handle := r.Group("/")
		{
			handle.POST("", f.Service.Del)
		}
	}
}
