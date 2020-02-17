package routes

import (
	Controllers "github.com/Stylll/pems-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	InitializeRoutes(r)

	return r
}

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", Controllers.Default)

	v1 := r.Group("api/v1")
	{
		v1.GET("/", Controllers.WelcomeAPI)
	}
}
