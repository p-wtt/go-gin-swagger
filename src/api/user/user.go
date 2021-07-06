package user

import(
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

// GetUserInfo godoc
// @Summary get user name
// @Description get user name, id
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true1 "User name"
// @Success 200 {object} UserInfo
// @Router /user/{name} [get]
func GetUserInfo(c *gin.Context) {
	name := c.Param("name")
	message := name + " is very handsome!"
	userInfo := UserInfo{1, message}

	// userInfo := model.User{ID: 1, Name: name}

	c.JSON(http.StatusOK, gin.H{"message": userInfo})
}