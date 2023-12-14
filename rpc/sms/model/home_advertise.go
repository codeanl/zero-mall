package model

import (
	"fmt"
	"gorm.io/gorm"
	"simple_mall_new/rpc/sms/sms"
)

type (
	HomeAdvertiseModel interface {
		AddHomeAdvertise(coupon *HomeAdvertise) (err error)
		UpdateHomeAdvertise(id int64, coupon *HomeAdvertise) error
		DeleteHomeAdvertiseByIds(ids []int64) error
		GetHomeAdvertiseList(in *sms.HomeAdvertiseListReq) ([]*HomeAdvertise, int64, error)
		SaveOrUpdateHomeAdvertise(id int64, req *HomeAdvertise) (err error)
	}
	defaultHomeAdvertiseModel struct {
		conn *gorm.DB
	}
	HomeAdvertise struct {
		gorm.Model
		Name   string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`  //名称
		Pic    string `json:"pic" gorm:"type:varchar(191);comment:图片地址;not null"` //图片地址
		Status string `json:"status" gorm:"type:varchar(191);comment:上下线状态"`      //上下线状态：0->下线；1->上线',
		Sort   int64  `json:"sort" gorm:"type:bigint;comment:排序"`                 //排序
		Url    string `json:"url" gorm:"type:varchar(191);comment:链接地址"`          //链接地址
		Note   string `json:"note" gorm:"type:varchar(191);comment:备注"`           //备注
	}
)

func NewHomeAdvertiseModel(conn *gorm.DB) HomeAdvertiseModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&HomeAdvertise{})
	return &defaultHomeAdvertiseModel{
		conn: conn,
	}
}

func (m *defaultHomeAdvertiseModel) AddHomeAdvertise(coupon *HomeAdvertise) (err error) {
	return m.conn.Model(&HomeAdvertise{}).Create(coupon).Error
}

func (m *defaultHomeAdvertiseModel) UpdateHomeAdvertise(id int64, coupon *HomeAdvertise) error {
	err := m.conn.Model(&HomeAdvertise{}).Where("id=?", id).Updates(coupon).Error
	return err
}

func (m *defaultHomeAdvertiseModel) DeleteHomeAdvertiseByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&HomeAdvertise{}).Error
	return err
}

func (m *defaultHomeAdvertiseModel) GetHomeAdvertiseList(in *sms.HomeAdvertiseListReq) ([]*HomeAdvertise, int64, error) {
	var list []*HomeAdvertise
	db := m.conn.Model(&list).Order("sort ASC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Status != "" {
		db = db.Where("status = ?", in.Status)
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	if in.PageNum > 0 && in.PageSize > 0 {
		err = db.Offset(int((in.PageNum - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}

func (m *defaultHomeAdvertiseModel) SaveOrUpdateHomeAdvertise(id int64, req *HomeAdvertise) (err error) {
	if id > 0 {
		return m.conn.Model(&HomeAdvertise{}).Where("id=?", id).Updates(req).Error
	} else {
		return m.conn.Model(&HomeAdvertise{}).Create(req).Error
	}
}
