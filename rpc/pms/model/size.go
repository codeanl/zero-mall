package model

import (
	"gorm.io/gorm"
)

type (
	SizeModel interface {
		AddSize(Size *Size) (*Size, error)
		UpdateSize(id int64, role *Size) error
		DeleteSizeByIds(ids []int64) error
		//GetSizeList(in *pms.SizeListReq) ([]*Size, int64, error)
		GetSizeById(id int64) (info *Size, err error)
		DeleteSizeByProID(id int64) error
		GetSizeByTag(tag string) (info *Size, err error)
		DeleteSizeByTag(tag string) error
		DeleteSizeByID(id int64) error
		GetSizeList(id int64) ([]*Size, int64, error)
	}

	defaultSizeModel struct {
		conn *gorm.DB
	}
	Size struct {
		ProductID int64  `json:"product_id" gorm:"type:bigint;comment:商品id;not null"` //商品id
		Name      string `json:"name" gorm:"type:varchar(191);comment:规格名称;not null"` //规格名称
		Value     string `json:"value" gorm:"type:varchar(191);comment:规格;not null"`  //规格
	}
)

func NewSizeModel(conn *gorm.DB) SizeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Size{})
	return &defaultSizeModel{
		conn: conn,
	}
}
func (m *defaultSizeModel) AddSize(info *Size) (*Size, error) {
	err := m.conn.Model(&Size{}).Create(info).Error
	return info, err
}

func (m *defaultSizeModel) UpdateSize(id int64, role *Size) error {
	err := m.conn.Model(&Size{}).Where("id=?", id).Updates(role).Error
	return err
}
func (m *defaultSizeModel) GetSizeById(id int64) (info *Size, err error) {
	err = m.conn.Model(&Size{}).Where("id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultSizeModel) GetSizeByTag(tag string) (info *Size, err error) {
	err = m.conn.Model(&Size{}).Where("tag=?", tag).Find(&info).Error
	return info, err
}

// DeleteRoleByIds 删除角色
func (m *defaultSizeModel) DeleteSizeByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Size{}).Error
	return err
}

func (m *defaultSizeModel) DeleteSizeByProID(id int64) error {
	err := m.conn.Model(&Size{}).Where("product_id=?", id).Delete(&Size{}).Error
	return err
}
func (m *defaultSizeModel) DeleteSizeByTag(tag string) error {
	err := m.conn.Model(&Size{}).Where("tag=?", tag).Delete(&Size{}).Error
	return err
}
func (m *defaultSizeModel) DeleteSizeByID(id int64) error {
	err := m.conn.Model(&Size{}).Where("id=?", id).Delete(&Size{}).Error
	return err
}

func (m *defaultSizeModel) GetSizeList(id int64) ([]*Size, int64, error) {
	var list []*Size
	db := m.conn.Model(&Size{})
	if id != 0 {
		db = db.Where("product_id = ?", id)
	}
	err := db.Find(&list).Error
	if err != nil {
		return list, 0, err
	}
	total := int64(len(list))
	return list, total, err
}
