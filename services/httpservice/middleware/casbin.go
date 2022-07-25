package middleware

import (
	"net/http"

	"github.com/cpw0321/mammoth/logger"

	"github.com/cpw0321/mammoth/datasource/mysql"

	"github.com/gin-gonic/gin"

	"github.com/cpw0321/mammoth/internal"
	"github.com/cpw0321/mammoth/services/authservice"
)

// Casbin ...
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "请求头Authorization不能为空",
			})
			c.Abort()
			return
		}
		myClaims, err := internal.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// 获取请求的URI
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		as := authservice.New()
		role, err := as.GetRoleByUserID(myClaims.UserID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "获取权限信息失败",
			})
			c.Abort()
			return
		}
		sub := role.Name
		e, err := internal.NewCasbin(mysql.DB)
		if err != nil {
			logger.Log.Errorf("casbin NewCasbin is err:%v", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "权限加载失败",
			})
			c.Abort()
			return
		}
		// 判断策略中是否存在
		//e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
		isOK, err := e.Enforce(sub, obj, act)
		if err != nil {
			logger.Log.Errorf("casbin Enforce is err:%v", err)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "权限加载失败",
			})
			c.Abort()
			return
		}
		if !isOK {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusBadRequest,
				"message": "权限校验不通过",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
