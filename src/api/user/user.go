package user

import(
	"net/http"
	"github.com/gin-gonic/gin"
	common "github.com/p-wtt/go-gin-swagger/src/common"
)

type UserV1Router struct {
	group       *gin.RouterGroup
	// mDB         *gorm.DB
	// rDB         *gorm.DB
	// awsConf     common.AWSConfs
	// imageCenter *common.ContentsCenter
}

func NewUserV1Router(r common.Router, basePath string) UserV1Router{
	u := UserV1Router {
		group: r.Version.Group(basePath),
	}

	u.group.GET(":name", GetUserInfo)

	return u
};

type UserInfo struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

// GetUserInfo godoc
// @Summary 유저 정보 가져오기
// @Description 입력된 유저의 정보를 반환해주는 API
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Success 200 {object} UserInfo
// @Router /user/{name} [get]
func GetUserInfo(c *gin.Context) {
	name := c.Param("name")
	message := name + " is very handsome!"
	userInfo := UserInfo{1, message}

	// userInfo := model.User{ID: 1, Name: name}

	c.JSON(http.StatusOK, gin.H{"message": userInfo})
}