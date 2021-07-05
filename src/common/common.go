package common

import (
	"github.com/gin-gonic/gin"
)

//BasicRouter : Router interface,  included common functions.
type BasicRouter interface {
	// InitRouter(g *gin.RouterGroup)
	Descriptor() string
}
type Router struct {
	Engine *gin.Engine
	Port   string
}
