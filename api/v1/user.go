package v1

import (
	"github.com/cpw0321/mammoth/services/authservice"
	"github.com/cpw0321/mammoth/types/requests"
	"github.com/cpw0321/mammoth/types/responses"
	"github.com/gin-gonic/gin"
)

// IUserController ...
type IUserController interface {
	// Login 登录
	Login(c *gin.Context)
}

// UserController ...
type UserController struct {
	user authservice.IAuthservice
}

// NewUser ...
func NewUser() *UserController {
	return &UserController{
		user: authservice.New(),
	}
}

func (uc *UserController) Login(c *gin.Context) {
	var r requests.UserRequest
	err := c.BindJSON(&r)
	if err != nil {
		responses.Fail(c, 10000, err.Error())
		return
	}

	err = uc.user.Login(r.UserName, r.Password)
	if err != nil {
		responses.Fail(c, 10001, err.Error())
		return
	}
	responses.Success(c, "")
}
