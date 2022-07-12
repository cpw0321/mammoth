package authservice

import (
	"errors"
	"fmt"

	"github.com/cpw0321/mammoth/common"

	authresp "github.com/cpw0321/mammoth/types/responses/authservice"

	"github.com/cpw0321/mammoth/types/requests/authservice"

	"github.com/cpw0321/mammoth/datasource/mysql"
	"github.com/cpw0321/mammoth/services/authservice/models"
	"github.com/cpw0321/mammoth/utils"
	"gorm.io/gorm"
)

// IAuthservice ...
type IAuthservice interface {
	// Register 用户注册
	Register(userName string, password string) error
	// Login 用户登录
	Login(userName string, password string) (*models.User, error)
	// GetUserList 获取用户列表
	GetUserList(r authservice.UserListRequest) (*authresp.ListOfUserResponseBody, error)
	// GetRole 获取角色信息
	GetRole(name string) (*models.Role, error)
	// GetRoleByUserID 通过用户id获取角色信息
	GetRoleByUserID(userID uint) (*models.Role, error)
	// CreateRole 创建角色
	CreateRole(name string) error
	// CreateUserRole 用户绑定角色
	CreateUserRole(userID uint, roleID uint) error
	// GetRoleList 获取角色列表
	GetRoleList(r authservice.RoleListRequest) (*authresp.ListOfRoleResponseBody, error)
}

var _ IAuthservice = (*Authservice)(nil)

// Authservice ...
type Authservice struct {
	db *gorm.DB
}

// New ...
func New() *Authservice {
	return &Authservice{
		db: mysql.DB,
	}
}

func (as *Authservice) Register(userName string, password string) error {
	var user models.User
	err := as.db.Model(&models.User{}).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user.UserName = userName
			user.Password = utils.MD5(password)
			if err := as.db.Model(&models.User{}).Save(&user).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return errors.New("用户已存在")
}

func (as *Authservice) Login(userName string, password string) (*models.User, error) {
	var user models.User
	err := as.db.Model(&models.User{}).Where("user_name = ? and password = ?", userName, utils.MD5(password)).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (as *Authservice) GetUserList(r authservice.UserListRequest) (*authresp.ListOfUserResponseBody, error) {
	var page, pageSize int
	if r.Page == common.ZERO {
		page = common.PAGE
	}
	if r.PageSize == common.ZERO {
		pageSize = common.PAGE_SIZE
	}
	tx := as.db.Model(models.User{})
	var count int64
	err := tx.Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	if r.OrderField != "" && r.OrderType != "" {
		tx.Order(fmt.Sprintf("%s %s", r.OrderField, r.OrderType))
	} else {
		tx.Order("created_at desc")
	}
	if r.UserName != "" {
		tx.Where("user_name like ?", "%"+r.UserName+"%")
	}
	var users []models.User
	err = tx.Limit(r.PageSize).Offset((r.Page - 1) * r.PageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}
	var items []authresp.UserBody
	var item authresp.UserBody
	for _, v := range users {
		err := utils.CopyStruct(v, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	res := &authresp.ListOfUserResponseBody{
		Total:    count,
		Page:     page,
		PageSize: pageSize,
		List:     items,
	}
	return res, nil
}

func (as *Authservice) GetRoleByUserID(userID uint) (*models.Role, error) {
	var userRole models.UserRole
	err := as.db.Model(&models.Role{}).Where("user_id = ? ", userID).First(&userRole).Error
	if err != nil {
		return nil, err
	}
	var role models.Role
	err = as.db.Model(&models.Role{}).Where("id = ? ", userRole.RoleId).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (as *Authservice) GetRole(name string) (*models.Role, error) {
	var role models.Role
	err := as.db.Model(&models.Role{}).Where("name = ? ", name).First(&role).Error
	return &role, err
}

func (as *Authservice) CreateRole(name string) error {
	var role models.Role
	err := as.db.Model(&models.Role{}).Where("name = ?", name).First(&role).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			role.Name = name
			if err := as.db.Model(&models.Role{}).Save(&role).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return errors.New("角色已存在")
}

func (as *Authservice) CreateUserRole(userID uint, roleID uint) error {
	var userRole models.UserRole
	userRole.UserId = userID
	userRole.RoleId = roleID
	err := as.db.Model(&models.UserRole{}).Save(&userRole).Error
	if err != nil {
		return err
	}
	return nil
}

func (as *Authservice) GetRoleList(r authservice.RoleListRequest) (*authresp.ListOfRoleResponseBody, error) {
	var page, pageSize int
	if r.Page == common.ZERO {
		page = common.PAGE
	}
	if r.PageSize == common.ZERO {
		pageSize = common.PAGE_SIZE
	}
	tx := as.db.Model(models.Role{})
	var count int64
	err := tx.Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	if r.OrderField != "" && r.OrderType != "" {
		tx.Order(fmt.Sprintf("%s %s", r.OrderField, r.OrderType))
	} else {
		tx.Order("created_at desc")
	}
	if r.Name != "" {
		tx.Where("name like ?", "%"+r.Name+"%")
	}
	var roles []models.Role
	err = tx.Limit(r.PageSize).Offset((r.Page - 1) * r.PageSize).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	var items []authresp.RoleBody
	var item authresp.RoleBody
	for _, v := range roles {
		err := utils.CopyStruct(v, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	res := &authresp.ListOfRoleResponseBody{
		Total:    count,
		Page:     page,
		PageSize: pageSize,
		List:     items,
	}
	return res, nil
}
