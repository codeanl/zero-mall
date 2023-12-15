package svc

import (
	"simple_mall_new/rpc/pms/internal/config"
	"simple_mall_new/rpc/pms/model"
	"simple_mall_new/utils"
)

type ServiceContext struct {
	Config        config.Config
	CategoryModel model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := utils.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		CategoryModel: model.NewCategoryModel(conn),
	}
}
