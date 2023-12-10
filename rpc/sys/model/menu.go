package model

import (
	"fmt"
	"gorm.io/gorm"
)

type (
	MenuModel interface {
		GetMenuByName(name string) (menu *Menu, exist bool, err error)
		SaveOrUpdateMenu(id int64, req *Menu) (err error)
		GetMenuList() ([]Menu, int64, error)
		DeleteMenuByIds(ids []int64) error
	}

	defaultMenuModel struct {
		conn *gorm.DB
	}
	Menu struct {
		gorm.Model
		Name     string `json:"name" gorm:"type:varchar(191);comment:菜单名称;not null"`  //菜单名称
		ParentID int64  `json:"parent_id" gorm:"type:bigint;comment:父菜单ID;not null;"` //父菜单ID
		Url      string `json:"url" gorm:"type:varchar(191);comment:路径;not null"`     //路径
		Icon     string `json:"icon" gorm:"type:varchar(191);comment:菜单图标;not null"`  //菜单图标
		Type     string `json:"type" gorm:"type:varchar(191);comment:类型;not null"`    //类型  1：目录  2：菜单   3：按钮
		OrderNum int64  `json:"order_num" gorm:"type:bigint;comment:排序;not null"`     //排序
		TAG      string `json:"tag" gorm:"type:varchar(191);comment:标识;not null"`     //标识
	}
)

func NewMenuModel(conn *gorm.DB) MenuModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Menu{})
	return &defaultMenuModel{
		conn: conn,
	}
}

func (m *defaultMenuModel) GetMenuList() ([]Menu, int64, error) {
	var list []Menu
	db := m.conn.Model(&Menu{}).Order("order_num ASC")
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Find(&list).Error
	return list, total, err
}
func (m *defaultMenuModel) DeleteMenuByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Menu{}).Error
	return err
}

// GetMenuByName  根据用户名查询信息｜查询是或存在
func (m *defaultMenuModel) GetMenuByName(name string) (menu *Menu, exist bool, err error) {
	var count int64
	err = m.conn.Model(&Menu{}).Where("name=?", name).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&Menu{}).Where("name=?", name).First(&menu).Error
	if err != nil {
		return nil, false, err
	}
	return menu, true, nil
}

func (m *defaultMenuModel) SaveOrUpdateMenu(id int64, req *Menu) (err error) {
	fmt.Println(id)
	if id > 0 {
		return m.conn.Model(&Menu{}).Where("id=?", id).Updates(req).Error
	} else {
		return m.conn.Model(&Menu{}).Create(req).Error
	}
}
