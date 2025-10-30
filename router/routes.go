package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	// docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("opening", handler.ShowOpeningHandler)
		v1.POST("opening", handler.ShowOpeningHandler)
		v1.DELETE("opening", handler.ShowOpeningHandler)
		v1.PUT("opening", handler.ShowOpeningHandler)
		v1.GET("openings", handler.ShowOpeningHandler)
	}
	// Initialize Swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
