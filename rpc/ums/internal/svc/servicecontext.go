package svc

import (
	"simple_mall_new/rpc/ums/internal/config"
	"simple_mall_new/rpc/ums/model"
	"simple_mall_new/utils"
)

type ServiceContext struct {
	Config      config.Config
	MemberModel model.MemberModel
	//MemberLoginLogModel model.MemberLoginLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := utils.InitMysql(mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		MemberModel: model.NewMemberModel(conn),
		//MemberLoginLogModel: model.NewMemberLoginLogModel(conn),
	}
}
