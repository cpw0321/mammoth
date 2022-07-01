package middleware

import "C"
import (
	"net/http"

	"github.com/cpw0321/mammoth/datasource/mysql"

	"github.com/gin-gonic/gin"

	"github.com/cpw0321/mammoth/internal"
	"github.com/cpw0321/mammoth/services/authservice"
)

// Casbin ...
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*internal.MyClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		as := authservice.New()
		role, err := as.GetRoleByUserID(waitUse.UserID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "获取权限信息失败",
			})
			return
		}
		sub := role.Name
		e, err := internal.NewCasbin(mysql.DB)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "权限加载失败",
			})
			return
		}
		// 判断策略中是否存在
		//e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
		isOK, err := e.Enforce(sub, obj, act)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "权限加载失败",
			})
			return
		}
		if !isOK {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "权限校验不通过",
			})
			return
		}
		c.Next()
	}
}
