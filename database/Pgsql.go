package database

import (
	"HertzDemo/configs"
	"database/sql"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DbInstance = new(gorm.DB)

var err error

var sqlDB  = new(sql.DB)

func InitPgsql(pgsql *configs.Pgsql) error {
	dsn := "host=" + pgsql.Host +
		" user=" + pgsql.Username +
		" password=" + pgsql.Password +
		" dbname=" + pgsql.DbName +
		" port=" + pgsql.Port +
		" sslmode=disable TimeZone=Asia/Shanghai"
	DbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		//DryRun: true,
	})
	if err != nil {
		return err
	}
	sqlDB, _ = DbInstance.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func Close()  {
	hlog.Info("数据库链接关闭了~~~~~~~~~~~~~~~~~~~~")
	sqlDB.Close()
}
