package model

import (
	"gorm.io/gorm"
)

type (
	UserModel interface {
		GetUserByUsername(username string) (user *User, exist bool, err error)
	}
	defaultUserModel struct {
		conn *gorm.DB
	}
	User struct {
		gorm.Model
		Username string `json:"username" gorm:"type:varchar(191);comment:用户名;not null"` //用户名
		Phone    string `json:"phone" gorm:"type:varchar(191);comment:手机号;not null"`    //手机号
		Nickname string `json:"nickname" gorm:"type:varchar(191);comment:昵称;not null"`  //昵称
		Password string `json:"password" gorm:"type:varchar(191);comment:密码;not null"`  //密码
		Gender   string `json:"gender" gorm:"type:varchar(255);comment:性别;"`            //性别  0--保密 1--男  2--女
		Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:用户头像"`           //头像
		Email    string `json:"email" gorm:"type:varchar(255);comment:邮箱"`              //邮箱
		Status   string `json:"status" gorm:"type:varchar(255);comment:状态;"`            //状态  1--正常 0--禁用 默认正常
	}
)

func NewUserModel(conn *gorm.DB) UserModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&User{})
	return &defaultUserModel{
		conn: conn,
	}
}

// GetUserByUsername  根据用户名查询用户信息｜查询用户是或存在
func (m *defaultUserModel) GetUserByUsername(username string) (user *User, exist bool, err error) {
	var count int64
	err = m.conn.Model(&User{}).Where("username=?", username).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&User{}).Where("username=?", username).First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}
