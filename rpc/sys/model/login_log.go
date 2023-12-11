package model

import (
	"gorm.io/gorm"
	"simple_mall_new/rpc/sys/sys"
)

type (
	LoginLogModel interface {
		AddLoginLog(loginLog *LoginLog) (err error)
		DeleteLoginLogByIds(ids []int64) error
		GetLoginLogList(in *sys.LoginLogListReq) ([]*LoginLog, int64, error)
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

func (m *defaultLoginLogModel) DeleteLoginLogByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&LoginLog{}).Error
	return err
}

func (m *defaultLoginLogModel) GetLoginLogList(in *sys.LoginLogListReq) ([]*LoginLog, int64, error) {
	var list []*LoginLog
	db := m.conn.Model(&LoginLog{}).Order("created_at DESC")

	if in.UserId != 0 {
		db = db.Where("user_id = ?", in.UserId)
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
