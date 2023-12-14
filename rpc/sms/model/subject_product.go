package model

import (
	"gorm.io/gorm"
	"simple_mall_new/rpc/sms/sms"
)

type (
	SubjectProductModel interface {
		AddSubjectProduct(info *SubjectProduct) (err error)
		UpdateSubjectProduct(id int64, info *SubjectProduct) error
		DeleteSubjectProductByIds(ids []int64) error
		GetSubjectProductList(in *sms.SubjectProductListReq) ([]*SubjectProduct, int64, error)
		////GetUserByID(id int64) (user *User, err error)
		ExistSubjectProductById(id int64) (info *SubjectProduct, exist bool, err error)
		GetSubjectProductExist(SubjectID, ProductID int64) (exist bool, err error)
	}
	defaultSubjectProductModel struct {
		conn *gorm.DB
	}
	SubjectProduct struct {
		gorm.Model
		SubjectID int64 `json:"subject_id" gorm:"type:bigint;comment:id;not null"`
		ProductID int64 `json:"product_id" gorm:"type:bigint;comment:id;not null"`
		Sort      int64 `json:"sort" gorm:"type:bigint;comment:排序;not null"`
	}
)

func NewSubjectProductModel(conn *gorm.DB) SubjectProductModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&SubjectProduct{})
	return &defaultSubjectProductModel{
		conn: conn,
	}
}

func (m *defaultSubjectProductModel) AddSubjectProduct(info *SubjectProduct) (err error) {
	return m.conn.Model(&SubjectProduct{}).Create(info).Error
}
func (m *defaultSubjectProductModel) ExistSubjectProductById(id int64) (info *SubjectProduct, exist bool, err error) {
	var count int64
	err = m.conn.Model(&SubjectProduct{}).Where("product_id=?", id).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&SubjectProduct{}).Where("product_id=?", id).First(&info).Error
	if err != nil {
		return nil, false, err
	}
	return info, true, nil
}

func (m *defaultSubjectProductModel) UpdateSubjectProduct(id int64, info *SubjectProduct) error {
	err := m.conn.Model(&SubjectProduct{}).Where("id=?", id).Updates(info).Error
	return err
}

func (m *defaultSubjectProductModel) DeleteSubjectProductByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&SubjectProduct{}).Error
	return err
}

// //GetUserList 获取用户列表
func (m *defaultSubjectProductModel) GetSubjectProductList(in *sms.SubjectProductListReq) ([]*SubjectProduct, int64, error) {
	var list []*SubjectProduct
	db := m.conn.Model(&list).Order("sort ASC")

	if in.SubjectId != 0 {
		db = db.Where("subject_id =?", in.SubjectId)
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
func (m *defaultSubjectProductModel) GetSubjectProductExist(SubjectID, ProductID int64) (exist bool, err error) {
	var count int64
	err = m.conn.Model(&SubjectProduct{}).Where("subject_id=? && product_id=?", SubjectID, ProductID).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
