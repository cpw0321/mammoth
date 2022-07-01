package v1

import (
	"net/http"

	"github.com/cpw0321/mammoth/internal"
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

func (uc *UserController) Login(c *internal.ServiceContext) {
	var r requests.UserRequest
	err := c.Context.ShouldBind(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	err = uc.user.Login(r.UserName, r.Password)
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.Success(c, "")
}

func (uc *UserController) Register(c *internal.ServiceContext) {
	var r requests.UserRequest
	err := c.Context.ShouldBindJSON(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	err = uc.user.Login(r.UserName, r.Password)
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.Success(c, "")
}
