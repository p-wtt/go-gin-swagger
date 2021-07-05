package routesV1

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/happierall/l"
	common "github.com/p-wtt/go-gin-swagger/src/common"

	"github.com/p-wtt/go-gin-swagger/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	whiteList = []string{
		"http://localhost:3000",       // localhost
		"http://182.224.216.110:3000", // hashlike sangmin
	}

	allowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	allowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "uid"}
)

func NewRoute() common.Router {
	return common.Router{
		Engine: gin.Default(),
		Port:   "9000",
	}
}

func initialize(r common.Router) bool {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = whiteList
	corsConfig.AllowMethods = allowMethods
	corsConfig.AllowMethods = allowHeaders
	corsConfig.AllowCredentials = true

	return true
}

func Start(r common.Router) {
	if _b := initialize(r); !_b {
		l.Error("Fail to intializing router.")
		return
	}

	// var RouterList []common.BasicRouter
	// RouterList = append(RouterList, swagger.NewSwaggerV1Router(r))
	// RouterList = append(RouterList, auth.NewAuthV1Router(r, "/auth"))

	// programatically set swagger info
	docs.SwaggerInfo.Title = "gin swagger test"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger.!!!!!!"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = "/v2"
	r.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	l.Log(r.Port)
	r.Engine.Run(fmt.Sprintf(":%s", r.Port))
}
