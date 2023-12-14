package model

import (
	"fmt"
	"gorm.io/gorm"
	"simple_mall_new/rpc/ums/ums"
)

type (
	MemberModel interface {
		GetMemberByPhone(phone string) (member *Member, exist bool, err error)
		AddMember(Member *Member) (err error)
		UpdateMember(id int64, Member *Member) error
		DeleteByIds(ids []int64) error
		GetMemberList(in *ums.MemberListReq) ([]*Member, int64, error)
		GetMemberByID(id int64) (Member *Member, err error)
		SaveOrUpdateMember(id int64, req *Member) (err error)
	}
	defaultMemberModel struct {
		conn *gorm.DB
	}
	Member struct {
		gorm.Model
		Phone    string `json:"phone" gorm:"type:varchar(191);comment:手机号;not null"`   //手机号
		Password string `json:"password" gorm:"type:varchar(191);comment:密码;not null"` //密码
		Nickname string `json:"nickname" gorm:"type:varchar(191);comment:昵称;not null"` //昵称
		Status   string `json:"status" gorm:"type:varchar(255);comment:状态;"`           //状态  0--正常 1--禁用 默认正常
		Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:用户头像"`          //头像
		Gender   string `json:"gender" gorm:"type:varchar(255);comment:性别;"`           //性别  0--保密 1--男  2--女
		Email    string `json:"email" gorm:"type:varchar(255);comment:邮箱"`             //邮箱
	}
)

func NewMemberModel(conn *gorm.DB) MemberModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Member{})
	return &defaultMemberModel{
		conn: conn,
	}
}
func (m *defaultMemberModel) GetMemberByID(id int64) (info *Member, err error) {
	err = m.conn.Model(&Member{}).Where("id=?", id).First(&info).Error
	return info, err
}

// GetMemberByPhone 查询用户是或存在
func (m *defaultMemberModel) GetMemberByPhone(phone string) (member *Member, exist bool, err error) {
	var count int64
	err = m.conn.Model(&Member{}).Where("phone=?", phone).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&Member{}).Where("phone=?", phone).First(&member).Error
	if err != nil {
		return nil, false, err
	}
	return member, true, nil
}

// AddMember 创建用户
func (m *defaultMemberModel) AddMember(info *Member) (err error) {
	return m.conn.Model(&Member{}).Create(info).Error
}

// UpdateMember 更新用户
func (m *defaultMemberModel) UpdateMember(id int64, info *Member) error {
	err := m.conn.Model(&Member{}).Where("id=?", id).Updates(info).Error
	return err
}

// DeleteByIds 删除用户
func (m *defaultMemberModel) DeleteByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Member{}).Error
	return err
}

// GetMemberList 获取用户列表
func (m *defaultMemberModel) GetMemberList(in *ums.MemberListReq) ([]*Member, int64, error) {
	var list []*Member
	db := m.conn.Model(&list).Order("created_at DESC")
	if in.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", in.Phone))
	}
	if in.Status != "" {
		db = db.Where("status= ?", in.Status)
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
func (m *defaultMemberModel) SaveOrUpdateMember(id int64, req *Member) (err error) {
	if id > 0 {
		return m.conn.Model(&Member{}).Where("id=?", id).Updates(req).Error
	} else {
		return m.conn.Model(&Member{}).Create(req).Error
	}
}
