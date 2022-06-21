package authservice

import (
	"github.com/cpw0321/mammoth/datasource/mysql"
	"gorm.io/gorm"
)

// IAuthservice ...
type IAuthservice interface {
	Login(userName string, password string) error
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
func (as *Authservice) Login(userName string, password string) error {
	var user User
	err := as.db.Model(&User{}).Where("userName = ? and password = ?", userName, password).First(&user).Error
	return err
}
