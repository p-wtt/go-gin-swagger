package auth

import(
	"net/http"
	
	"github.com/gin-gonic/gin"
)

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