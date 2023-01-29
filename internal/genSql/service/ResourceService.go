package service

import (
	"HertzDemo/internal/genSql/handler/req"
	"HertzDemo/internal/genSql/model"
	"HertzDemo/internal/genSql/repository"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"strings"
)

const (
	FeiMai = 1
)

func LastOneByType(Type string) *model.Resource {
	resource := repository.LastOneByType(Type)
	return resource
}

func LastOrderNoByParentId(ParentId int) int {
	orderNo := repository.LastOrderNoByParentId(ParentId)
	return orderNo
}

func InsertResource(req req.InsertReq) int {
	//处理code
	code := req.Type
	lastResource := LastOneByType(strings.ToUpper(req.Type))
	var codeNum int
	if lastResource != nil {
		dbCodeNum, err := strconv.Atoi(lastResource.Code[1:])
		if err != nil {
			hlog.Error("参数异常！")
		}
		codeNum = dbCodeNum
	}
	codeNum += 1
	code += fmt.Sprintf("%07d", codeNum)
	hlog.Info(code)
	// 处理order_no
	orderNo := LastOrderNoByParentId(req.ParentId)

	var resource = new(model.Resource)
	resource.OrderNo = orderNo
	resource.Code = code
	resource.Name = req.Name
	resource.Type = req.Type
	resource.Visible = req.Visible
	resource.ParentId = req.ParentId
	resource.Path = req.Path
	resource.Resource = req.Resource
	resource.Button = req.Button
	resource.System = FeiMai

	return repository.InsertResource(resource)
}
