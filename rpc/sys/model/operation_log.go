package model

import (
	"fmt"
	"gorm.io/gorm"
	"simple_mall_new/rpc/sys/sys"
)

type (
	OperationLogModel interface {
		AddOperationLog(OperationLog *OperationLog) (err error)
		GetOperationLogList(in *sys.OperationLogListReq) ([]*OperationLog, int64, error)
		DeleteOperationLogByIds(ids []int64) error
	}

	defaultOperationLogModel struct {
		conn *gorm.DB
	}
	OperationLog struct {
		gorm.Model
		UserID    int64  `json:"user_id" gorm:"type:bigint;comment:用户名;not null"`          //用户名
		Operation string `json:"operation" gorm:"type:varchar(191);comment:用户操作;not null"` //用户操作
		Method    string `json:"method" gorm:"type:varchar(1000);comment:请求方法;not null"`   //请求方法
		Params    string `json:"params" gorm:"type:varchar(1000);comment:请求参数;not null"`   //请求参数
		Time      int64  `json:"time" gorm:"type:varchar(191);comment:执行时长(毫秒);not null"`  //执行时长(毫秒)
		IP        string `json:"ip" gorm:"type:varchar(191);comment:IP地址;not null"`        //IP地址
	}
)

func NewOperationLogModel(conn *gorm.DB) OperationLogModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&OperationLog{})
	return &defaultOperationLogModel{
		conn: conn,
	}
}
func (m *defaultOperationLogModel) AddOperationLog(log *OperationLog) (err error) {
	return m.conn.Model(&OperationLog{}).Create(log).Error
}
func (m *defaultOperationLogModel) GetOperationLogList(in *sys.OperationLogListReq) ([]*OperationLog, int64, error) {
	var list []*OperationLog
	db := m.conn.Model(&OperationLog{}).Order("created_at DESC")
	if in.Method != "" {
		db = db.Where("method LIKE ?", fmt.Sprintf("%%%s%%", in.Method))
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
func (m *defaultOperationLogModel) DeleteOperationLogByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&OperationLog{}).Error
	return err
}
