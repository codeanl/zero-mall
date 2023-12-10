package model

import (
	"gorm.io/gorm"
)

type (
	LoginLogModel interface {
		AddLoginLog(loginLog *LoginLog) (err error)
	}

	defaultLoginLogModel struct {
		conn *gorm.DB
	}
	LoginLog struct {
		gorm.Model
		UserID  int64  `json:"user_id" gorm:"type:bigint;comment:用户名;not null"`           //用户id
		IP      string `json:"ip" gorm:"type:varchar(191);comment:IP地址;not null"`         //IP地址
		Address string `json:"address" gorm:"type:varchar(191);comment:ip归属地地址;not null"` //ip归属地地址
	}
)

func NewLoginLogModel(conn *gorm.DB) LoginLogModel {
	conn.AutoMigrate(&LoginLog{})
	return &defaultLoginLogModel{
		conn: conn,
	}
}
func (m *defaultLoginLogModel) AddLoginLog(loginLog *LoginLog) (err error) {
	return m.conn.Model(&LoginLog{}).Create(loginLog).Error
}
