package model

import (
	"fmt"
	"gorm.io/gorm"
	"simple_mall_new/rpc/pms/pms"
	"time"
)

type (
	MerchantsApplyModel interface {
		AddMerchantsApply(info *MerchantsApply) (*MerchantsApply, error)
		UpdateMerchantsApply(id int64, role *MerchantsApply) error
		DeleteMerchantsApplyBySpuID(id int64) error
		GetImgtByProducID(id int64) ([]MerchantsApply, error)
		DeleteMerchantsApplyByIds(ids []int64) error
		GetMerchantsApplyList(in *pms.MerchantsApplyListReq) ([]*MerchantsApply, int64, error)
		GetMerchantsApplyByID(id int64) (MerchantsApply, error)
		GetMerchantsApplyByPhone(phone string) (cate *MerchantsApply, exist bool, err error)
	}
	defaultMerchantsApplyModel struct {
		conn *gorm.DB
	}
	MerchantsApply struct {
		gorm.Model
		PrincipalName  string    `json:"principal_name" gorm:"type:varchar(225);comment:负责人名字;not null"`    //负责人名字
		PrincipalPhone string    `json:"principal_phone" gorm:"type:varchar(225);comment:负责人联系电话;not null"` //负责人联系电话
		IDCardFront    string    `json:"id_card_front" gorm:"type:varchar(225);comment:身份证正面;not null"`     //身份证正面
		IDCardReverse  string    `json:"id_card_reverse" gorm:"type:varchar(225);comment:身份证反面;not null"`   //身份证反面
		Name           string    `json:"name" gorm:"type:varchar(225);comment:身份证反面;not null"`              //店
		Address        string    `json:"address" gorm:"type:varchar(225);comment:地址;not null"`              //地址
		Pic            string    `json:"pic" gorm:"type:varchar(225);comment:地址;not null"`                  //店面图片
		Type           string    `json:"type" gorm:"type:varchar(225);comment:类型;not null"`                 //类型：0->商家 1->自提点
		Status         string    `json:"status" gorm:"type:varchar(225);comment:审核状态;not null"`             //审核状态 0->申请中 1->申请成功 2->申请失败
		Auditor        string    `json:"auditor" gorm:"type:varchar(225);comment:审核人;not null"`             //审核人
		ApprovalTime   time.Time `json:"approval_time" gorm:"type:datetime;comment:审核时间;not null"`          //审核时间
		Remark         string    `json:"remark" gorm:"type:varchar(225);comment:备注;not null"`               //备注
		AdminRemark    string    `json:"admin_remark" gorm:"type:varchar(225);comment:管理员备注;not null"`      //备注
	}
)

func NewMerchantsApplyModel(conn *gorm.DB) MerchantsApplyModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&MerchantsApply{})
	return &defaultMerchantsApplyModel{
		conn: conn,
	}
}
func (m *defaultMerchantsApplyModel) AddMerchantsApply(info *MerchantsApply) (*MerchantsApply, error) {
	err := m.conn.Model(&MerchantsApply{}).Create(info).Error
	return info, err
}

// UpdateRole 修改角色
func (m *defaultMerchantsApplyModel) UpdateMerchantsApply(id int64, role *MerchantsApply) error {
	err := m.conn.Model(&MerchantsApply{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultMerchantsApplyModel) DeleteMerchantsApplyBySpuID(id int64) error {
	err := m.conn.Model(&MerchantsApply{}).Where("product_id=?", id).Delete(&MerchantsApply{}).Error
	return err
}
func (m *defaultMerchantsApplyModel) GetAttributeValueByProductID(id int64) ([]MerchantsApply, error) {
	var info []MerchantsApply
	err := m.conn.Model(&MerchantsApply{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultMerchantsApplyModel) GetImgtByProducID(id int64) ([]MerchantsApply, error) {
	var info []MerchantsApply
	err := m.conn.Model(&MerchantsApply{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultMerchantsApplyModel) DeleteMerchantsApplyByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&MerchantsApply{}).Error
	return err
}

func (m *defaultMerchantsApplyModel) GetMerchantsApplyList(in *pms.MerchantsApplyListReq) ([]*MerchantsApply, int64, error) {
	var list []*MerchantsApply
	db := m.conn.Model(&MerchantsApply{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Status != "" {
		db = db.Where("status = ?", in.Status)
	}
	if in.Type != "" {
		db = db.Where("type = ?", in.Type)
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
func (m *defaultMerchantsApplyModel) GetMerchantsApplyByID(id int64) (MerchantsApply, error) {
	var info MerchantsApply
	err := m.conn.Model(&MerchantsApply{}).Where("id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultMerchantsApplyModel) GetMerchantsApplyByPhone(phone string) (cate *MerchantsApply, exist bool, err error) {
	var count int64
	err = m.conn.Model(&MerchantsApply{}).Where("principal_phone=?", phone).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&MerchantsApply{}).Where("principal_phone=?", phone).First(&cate).Error
	if err != nil {
		return nil, false, err
	}
	return cate, true, nil
}
