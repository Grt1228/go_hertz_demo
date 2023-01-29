package handler

import (
	"HertzDemo/internal/genSql/handler/req"
	"HertzDemo/internal/genSql/service"
	"HertzDemo/internal/pkg"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InsertResource(ctx context.Context, c *app.RequestContext) {
	var req req.InsertReq
	if err := c.BindAndValidate(&req); err != nil {
		hlog.Error("新增资源解析参数错误！")
		//存储错误到RequestContext中
		c.Error(errors.NewPublic("新增资源解析参数错误！"))
		panic("新增资源解析参数错误！")
	}
	hlog.Infof("请求参数：%s", req)
	pkg.Ok(c, service.InsertResource(req))
}
