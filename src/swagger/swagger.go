//swagger test file
package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/p-wtt/go-gin-swagger/docs"
	common "github.com/p-wtt/go-gin-swagger/src/common"
)

//PaymentV21Router : Struct about "PaymentV21Router"
type SwaggerV1Router struct {
	group *gin.RouterGroup
}

type welcomeModel struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

//Descriptor : Have to implement. If not, cannot use BasicRouter interface.
func (s SwaggerV1Router) Descriptor() string {
	return "SwaggerV1Router"
}

func NewSwaggerV1Router(r common.Router, subDomain string) SwaggerV1Router {
	s := SwaggerV1Router{
		group: r.Engine.Group(subDomain),
	}

	s.group.GET("welcome/:name", welcomePathParam)

	return s
}

// Welcome godoc
// @Summary Summary를 적어 줍니다.
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /test/welcome/{name} [get]
// @Success 200 {object} welcomeModel
func welcomePathParam(c *gin.Context) {
	docs.SwaggerInfo.BasePath = "/test"
	name := c.Param("name")
	message := name + " is very handsome!"
	welcomeMessage := welcomeModel{1, message}

	// welcomeMessage := model.User{ID: 1, Name: name}

	c.JSON(200, gin.H{"message": welcomeMessage})
}
