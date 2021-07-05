package auth

import (
	"github.com/gin-gonic/gin"
	common "github.com/p-wtt/go-gin-swagger/src/common"
)

//PaymentV21Router : Struct about "PaymentV21Router"
type AuthV1Router struct {
	group *gin.RouterGroup
}

type welcomeModel struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

//Descriptor : Have to implement. If not, cannot use BasicRouter interface.
func (s AuthV1Router) Descriptor() string {
	return "AuthV1Router"
}

func NewAuthV1Router(r common.Router, subDomain string) AuthV1Router {
	a := AuthV1Router{
		group: r.Engine.Group(subDomain),
	}

	return a
}
