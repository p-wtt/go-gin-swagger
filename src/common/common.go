package common

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	Port   string
	Version *gin.RouterGroup
	// Db     DataBases
	// Env    *viper.Viper
}

// type DataBases struct {
// 	MasterDB *gorm.DB
// 	ReadDB   *gorm.DB
// }