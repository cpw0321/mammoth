package mysql

import (
	"fmt"

	"gorm.io/gorm/schema"

	"github.com/cpw0321/mammoth/config"
	"github.com/cpw0321/mammoth/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据连接
var DB *gorm.DB

// InitDB mysql连接实例
// 默认用的配置文件地址MysqlDsn
func InitDB() {
	c := config.Conf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.DBName, "Asia%2FShanghai")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭AutoMigrate自动创建数据库外键约束
		SkipDefaultTransaction:                   true, // 禁用默认事务, 获得大约 30%+ 性能提升
		//Logger:                                   logger.NewGormLogger(logger.Log),
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: c.Mysql.DBTablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Log.Errorf("open mysql is failed, err:", err)
	}

	DB = db
}
