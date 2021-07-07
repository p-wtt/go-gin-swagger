package routes

import (
	//SECTION swagger 
	"github.com/p-wtt/go-gin-swagger/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	//SECTION router section
	common "github.com/p-wtt/go-gin-swagger/src/common"
	auth "github.com/p-wtt/go-gin-swagger/src/api/auth"
	user "github.com/p-wtt/go-gin-swagger/src/api/user"

	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gin-contrib/cors"
)

var (
	whiteList = []string{
		"http://localhost:3000",       // localhost
		"http://182.224.216.110:3000", // sangmin
	}

	allowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	allowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "uid"}
)

//NOTE x project initialize
func Initialize() common.Router {
	portNumber := "9000"

	//CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = whiteList
	corsConfig.AllowMethods = allowMethods
	corsConfig.AllowHeaders = allowHeaders
	corsConfig.AllowCredentials = true

	return common.Router{
		Engine: gin.Default(),
		Port: portNumber,
	}
}

//NOTE x project run
func Start(r common.Router) {
	_r := endpointManagement(r)
	_r.Engine.Run(fmt.Sprintf(":%s", _r.Port))
}

//NOTE x swagger setting 
func endpointManagement(r common.Router) common.Router {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "gin swagger test"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger.!!!!!!"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = "/v1"

	r.Version = r.Engine.Group("/v1")

	//SECTION x endpoints by function
	user.NewUserV1Router(r, "/user")
	auth.NewAuthV1Router(r, "/auth")

	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}