package svc

import (
	"simple_mall_new/rpc/sms/internal/config"
	"simple_mall_new/rpc/sms/model"
	"simple_mall_new/utils"
)

type ServiceContext struct {
	Config             config.Config
	HomeAdvertiseModel model.HomeAdvertiseModel
	CouponModel        model.CouponModel
	SubjectModel       model.SubjectModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := utils.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		HomeAdvertiseModel: model.NewHomeAdvertiseModel(conn),
		CouponModel:        model.NewCouponModel(conn),
		SubjectModel:       model.NewSubjectModel(conn),
	}
}
