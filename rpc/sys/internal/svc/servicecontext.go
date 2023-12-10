package svc

import (
	"simple_mall_new/rpc/sys/internal/config"
	"simple_mall_new/rpc/sys/model"
	"simple_mall_new/utils"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	LoginLogModel model.LoginLogModel
	RoleModel     model.RoleModel
	UserRoleModel model.UserRoleModel
	MenuModel     model.MenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := utils.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(conn),
		LoginLogModel: model.NewLoginLogModel(conn),
		RoleModel:     model.NewRoleModel(conn),
		UserRoleModel: model.NewUserRoleModel(conn),
		MenuModel:     model.NewMenuModel(conn),
	}
}
