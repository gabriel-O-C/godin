package router

import (
	"github.com/gabriel-O-C/godin/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}

}
