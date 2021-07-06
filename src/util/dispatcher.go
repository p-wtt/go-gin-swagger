//NOTE x endpoint management file
package util

import (
	"github.com/p-wtt/go-gin-swagger/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	auth "github.com/p-wtt/go-gin-swagger/src/api/auth"
	user "github.com/p-wtt/go-gin-swagger/src/api/user"

	"github.com/gin-gonic/gin"
)

func EndpointManagement(r *gin.Engine, /*db*/) *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "gin swagger test"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger.!!!!!!"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = "/v1"

	v1 := r.Group("/v1")
	{
		gauth := v1.Group("/auth")
		{
			gauth.GET("healthCheck", auth.HealthCheckMethod)
		}
		guser := v1.Group("/user")
		{
			guser.GET(":name", user.GetUserInfo)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}