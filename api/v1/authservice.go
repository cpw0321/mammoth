package v1

import (
	"net/http"

	authresp "github.com/cpw0321/mammoth/types/responses/authservice"

	"github.com/cpw0321/mammoth/logger"

	authreq "github.com/cpw0321/mammoth/types/requests/authservice"

	"github.com/cpw0321/mammoth/internal"
	"github.com/cpw0321/mammoth/services/authservice"
	"github.com/cpw0321/mammoth/types/responses"
)

// IAuthController ...
type IAuthController interface {
	// Login 登录
	Login(c *internal.ServiceContext)
	// Register 注册
	Register(c *internal.ServiceContext)
	// UserList 获取用户列表
	UserList(c *internal.ServiceContext)
	// RoleCreate 新建角色
	RoleCreate(c *internal.ServiceContext)
	// UserRole 分配用户角色
	UserRole(c *internal.ServiceContext)
	// RoleList 获取角色列表
	RoleList(c *internal.ServiceContext)
}

// AuthController ...
type AuthController struct {
	auth authservice.IAuthservice
}

// NewAuth ...
func NewAuth() *AuthController {
	return &AuthController{
		auth: authservice.New(),
	}
}

func (as *AuthController) Login(c *internal.ServiceContext) {
	var r authreq.UserRequest
	err := c.Context.ShouldBindJSON(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	user, err := as.auth.Login(r.UserName, r.Password)
	if err != nil {
		logger.Log.Errorf("Login Login is failed, err:%v", err)
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := internal.CreateToken(user.ID)
	if err != nil {
		logger.Log.Errorf("Login CreateToken is failed, err:%v", err)
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	var res authresp.UserLoginResponseBody
	res.Token = token
	responses.Success(c, res)
}

func (as *AuthController) Register(c *internal.ServiceContext) {
	var r authreq.UserRequest
	err := c.Context.ShouldBindJSON(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	err = as.auth.Register(r.UserName, r.Password)
	if err != nil {
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.Success(c, "")
}

func (as *AuthController) UserList(c *internal.ServiceContext) {
	var r authreq.UserListRequest
	err := c.Context.ShouldBind(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	userList, err := as.auth.GetUserList(r)
	if err != nil {
		logger.Log.Errorf("UserList GetUserList is failed, err:%v", err)
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.Success(c, userList)
}

func (as *AuthController) UserRole(c *internal.ServiceContext) {
	var r authreq.UserRoleRequest
	err := c.Context.ShouldBindJSON(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	err = as.auth.CreateUserRole(r.UserID, r.RoleID)
	if err != nil {
		logger.Log.Errorf("UserRole CreateUserRole is failed, err:%v", err)
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.Success(c, "")
}

func (as *AuthController) RoleCreate(c *internal.ServiceContext) {
	var r authreq.RoleCreateRequest
	err := c.Context.ShouldBindJSON(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	err = as.auth.CreateRole(r.Name)
	if err != nil {
		logger.Log.Errorf("RoleCreate CreateRole is failed, err:%v", err)
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.Success(c, "")
}

func (as *AuthController) RoleList(c *internal.ServiceContext) {
	var r authreq.RoleListRequest
	err := c.Context.ShouldBind(&r)
	if err != nil {
		responses.Fail(c, http.StatusBadRequest, internal.Translate(err))
		return
	}

	res, err := as.auth.GetRoleList(r)
	if err != nil {
		logger.Log.Errorf("RoleList GetRoleList is failed, err:%v", err)
		responses.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.Success(c, res)
}
