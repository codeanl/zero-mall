package model

import (
	"fmt"
	"gorm.io/gorm"
	"simple_mall_new/rpc/sys/sys"
)

type (
	RoleModel interface {
		SaveOrUpdateRole(id int64, req *Role) (err error)
		DeleteRoleByIds(ids []int64) error
		GetRoleList(req *sys.RoleListReq) ([]*Role, int64, error)
		GetRoleByName(name string) (role *Role, exist bool, err error)
		GetRoleByUserID(userid int64) ([]Role, error)
	}

	defaultRoleModel struct {
		conn *gorm.DB
	}
	Role struct {
		gorm.Model
		Name     string `json:"name" gorm:"type:varchar(191);comment:角色名称;not null"`     //角色名称
		Remark   string `json:"remark" gorm:"type:varchar(191);comment:备注;not null"`     //备注
		CreateBy string `json:"create_by" gorm:"type:varchar(191);comment:创建人;not null"` //创建人
		UpdateBy string `json:"update_by" gorm:"type:varchar(191);comment:更新人;not null"` //更新人
	}
)

func NewRoleModel(conn *gorm.DB) RoleModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Role{})
	return &defaultRoleModel{
		conn: conn,
	}
}

// GetRoleByName  根据用户名查询用户信息｜查询用户是或存在
func (m *defaultRoleModel) GetRoleByName(name string) (role *Role, exist bool, err error) {
	var count int64
	err = m.conn.Model(&Role{}).Where("name=?", name).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&Role{}).Where("name=?", name).First(&role).Error
	if err != nil {
		return nil, false, err
	}
	return role, true, nil
}

// SaveOrUpdateRole
func (m *defaultRoleModel) SaveOrUpdateRole(id int64, req *Role) (err error) {
	fmt.Println(id)
	if id > 0 {
		return m.conn.Model(&Role{}).Where("id=?", id).Updates(req).Error
	} else {
		return m.conn.Model(&Role{}).Create(req).Error
	}
}

// DeleteRoleByIds 删除角色
func (m *defaultRoleModel) DeleteRoleByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Role{}).Error
	//todo 手动删除相关的 userrole 数据->gorm关联删除
	err = m.conn.Where("role_id IN (?)", ids).Delete(&UserRole{}).Error
	return err
}

// GetRoleList 获取列表
func (m *defaultRoleModel) GetRoleList(req *sys.RoleListReq) ([]*Role, int64, error) {
	var list []*Role
	db := m.conn.Model(&Role{}).Order("created_at DESC")
	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	if req.PageNum > 0 && req.PageSize > 0 {
		err = db.Offset(int((req.PageNum - 1) * req.PageSize)).Limit(int(req.PageSize)).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}

// GetRoleByUserID 通过用户id查询用户的所有角色
func (m *defaultRoleModel) GetRoleByUserID(userid int64) ([]Role, error) {
	var userrole []UserRole
	if err := m.conn.Model(&UserRole{}).Where("user_id=?", userid).Find(&userrole).Error; err != nil {
		return nil, err
	}
	var roles []Role
	for _, item := range userrole {
		var role Role
		if err := m.conn.Model(&Role{}).Where("id = ?", item.RoleID).First(&role).Error; err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
