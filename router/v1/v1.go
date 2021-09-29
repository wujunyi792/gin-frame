package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadApi(e *gin.Engine) {
	apiGroup := e.Group("/v1")
	{
		apiGroup.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
}
