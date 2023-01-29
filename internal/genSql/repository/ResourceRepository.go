package repository

import (
	"HertzDemo/database"
	"HertzDemo/internal/genSql/model"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

const (
	DefaultUserId       = -1
	DefaultOrgId        = -1
	DefaultStatus       = 1
	DefaultDeleteStatus = 0
)

func LastOneByType(Type string) *model.Resource {
	var resource = new(model.Resource)
	result := database.DbInstance.Where("type = ?", Type).Order("code desc").Limit(1).Find(&resource)
	if result.RowsAffected == 0 {
		return nil
	}
	return resource
}

func LastOrderNoByParentId(ParentId int) int {
	var resource = new(model.Resource)
	result := database.DbInstance.Where("parent_id = ?", ParentId).Order("order_no desc").Limit(1).Find(&resource)
	if result.RowsAffected == 0 {
		return 10
	}
	return resource.OrderNo + 10
}

func InsertResource(resource *model.Resource) int {
	resource.GenTime = time.Now()
	resource.ModifiedTime = time.Now()
	resource.GenUserId = DefaultUserId
	resource.DeleteStatus = DefaultDeleteStatus
	resource.Status = DefaultStatus
	resource.OrgId = DefaultOrgId
	sql := database.DbInstance.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(resource)
	})
	hlog.Info(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	hlog.Info(sql)
	result := database.DbInstance.Create(resource)
	if result.RowsAffected == 0 {
		hlog.Error("创建资源异常！")
		panic("创建资源异常！")
	}
	return resource.Id
}
