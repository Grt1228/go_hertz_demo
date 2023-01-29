package main

import (
	"HertzDemo/configs"
	"HertzDemo/database"
	"HertzDemo/internal/genSql/handler"
	"HertzDemo/middleware"
	"context"
	"flag"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"gorm.io/gorm"
)

func main() {
	var config string
	flag.StringVar(&config, "config", "./configs/config.yaml", "-config 配置文件路径")
	flag.Parse()
	hlog.Infof("配置文件目录：%s", config)
	// 读取配置文件
	if err := configs.ReadConfig(config); err != nil {
		hlog.Errorf("系统加载配置文件错误：%s", err)
	}

	h := server.Default(server.WithHostPorts("[::]:" + configs.Config.Server.Port))

	//初始化数据库链接
	if err := database.InitPgsql(configs.Config.Pgsql); err != nil {
		hlog.Errorf("系统初始化数据库链接错误：%s", err)
	}
	SQL := database.DbInstance.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Raw("SELECT 1 FROM USER ID = ? ;", 1)
	})
	hlog.Info(SQL)
	defer database.Close()
	hlog.Infof("%s启动配置端口：%s", configs.Config.Server.Name, configs.Config.Server.Port)
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.Use(middleware.GlobalErrorHandler())
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	RegisterResourceRoute(h)
	h.Spin()
}

func RegisterResourceRoute(h *server.Hertz) {
	h.POST("/resource", handler.InsertResource)
}
