package auth

import(
	"net/http"
	
	"github.com/gin-gonic/gin"
	common "github.com/p-wtt/go-gin-swagger/src/common"
)

type AuthV1Router struct {
	group       *gin.RouterGroup
	// mDB         *gorm.DB
	// rDB         *gorm.DB
	// awsConf     common.AWSConfs
	// imageCenter *common.ContentsCenter
}

func NewAuthV1Router(r common.Router, basePath string) AuthV1Router {
	a := AuthV1Router {
		group: r.Version.Group(basePath),
	}

	a.group.GET("healthCheck", HealthCheckMethod)

	return a
}

// HealthCheckMethod godoc
// @Summary health check for server
// @Description check server status
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/healthCheck [get]
func HealthCheckMethod(c *gin.Context) {
	c.JSON(http.StatusOK, "status OK!!")
}