package main

import (
	dispatcher "github.com/p-wtt/go-gin-swagger/src/util"
	"github.com/gin-gonic/gin"
	"fmt"

	"github.com/gin-contrib/cors"
)

type Agent struct {
	Router *gin.Engine
}


var (
	whiteList = []string{
		"http://localhost:3000",       // localhost
		"http://182.224.216.110:3000", // sangmin
	}

	allowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	allowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "uid"}
)

func main() {
	var webServerEvn struct {
		Port string
	}

	webServerEvn.Port = "9000"
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = whiteList
	corsConfig.AllowMethods = allowMethods
	corsConfig.AllowMethods = allowHeaders
	corsConfig.AllowCredentials = true

	r := gin.Default()

	ag := Agent{
		Router: gin.Default(),
	}
	ag.Router.Use(cors.New(corsConfig))
	r = dispatcher.EndpointManagement(ag.Router)
	r.Run(fmt.Sprintf(":%s", webServerEvn.Port))
}
