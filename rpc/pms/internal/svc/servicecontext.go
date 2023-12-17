package svc

import (
	"simple_mall_new/rpc/pms/internal/config"
	"simple_mall_new/rpc/pms/model"
	sysModel "simple_mall_new/rpc/sys/model"
	"simple_mall_new/utils"
)

type ServiceContext struct {
	Config                   config.Config
	CategoryModel            model.CategoryModel
	MerchantsApplyModel      model.MerchantsApplyModel
	MerchantsModel           model.MerchantsModel
	PlaceModel               model.PlaceModel
	UserModel                sysModel.UserModel
	ProductModel             model.ProductModel
	ProductImgModel          model.ProductImgModel
	ProductIntroduceImgModel model.ProductIntroduceImgModel
	SkuModel                 model.SkuModel
	SizeModel                model.SizeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := utils.InitMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                   c,
		CategoryModel:            model.NewCategoryModel(conn),
		MerchantsApplyModel:      model.NewMerchantsApplyModel(conn),
		MerchantsModel:           model.NewMerchantsModel(conn),
		PlaceModel:               model.NewPlaceModel(conn),
		UserModel:                sysModel.NewUserModel(conn),
		ProductModel:             model.NewProductModel(conn),
		ProductImgModel:          model.NewProductImgModel(conn),
		ProductIntroduceImgModel: model.NewProductIntroduceImgModel(conn),
		SkuModel:                 model.NewSkuModel(conn),
		SizeModel:                model.NewSizeModel(conn),
	}
}
