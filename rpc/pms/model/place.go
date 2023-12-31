package model

import (
	"fmt"
	"gorm.io/gorm"
	"simple_mall_new/rpc/pms/pms"
)

type (
	PlaceModel interface {
		AddPlace(role *Place) (err error)
		DeletePlaceByIds(ids []int64) error
		GetPlaceList(in *pms.PlaceListReq) ([]*Place, int64, error)
		UpdatePlace(id int64, role *Place) error
		//GetPlaceByIdOrUserID(in *sys.PlaceInfoReq) (place *Place, err error)
		GetPlaceByID(id int64) (user *Place, err error)
	}

	defaultPlaceModel struct {
		conn *gorm.DB
	}
	Place struct {
		gorm.Model
		UserID               int64  `json:"user_id" gorm:"type:bigint;comment:用户id;not null"`                 //用户id
		Name                 string `json:"name" gorm:"type:varchar(225);comment:名称;not null"`                //名称
		PrincipalName        string `json:"principal_name" gorm:"type:varchar(225);comment:负责人;not null"`     //负责人
		PrincipalPhone       string `json:"principal_phone" gorm:"type:varchar(225);comment:联系电话;not null"`   //联系电话
		Address              string `json:"address" gorm:"type:varchar(225);comment:地址;not null"`             //地址
		Pic                  string `json:"pic" gorm:"type:varchar(225);comment:图片;not null"`                 //图片
		MerchantApplysListID int64  `json:"merchant_applys_list_id" gorm:"type:bigint;comment:关联id;not null"` //关联id
	}
)

func NewPlaceModel(conn *gorm.DB) PlaceModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Place{})
	return &defaultPlaceModel{
		conn: conn,
	}
}

func (m *defaultPlaceModel) AddPlace(role *Place) (err error) {
	return m.conn.Model(&Place{}).Create(role).Error
}

func (m *defaultPlaceModel) UpdatePlace(id int64, role *Place) error {
	err := m.conn.Model(&Place{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultPlaceModel) DeletePlaceByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Place{}).Error
	return err
}

//func (m *defaultPlaceModel) GetPlaceByIdOrUserID(in *sys.PlaceInfoReq) (place *Place, err error) {
//	if in.Id != 0 {
//		err = m.conn.Model(&Place{}).Where("id=?", in.Id).First(&place).Error
//	}
//	if in.UserID != 0 {
//		err = m.conn.Model(&Place{}).Where("user_id=?", in.UserID).First(&place).Error
//	}
//	return place, err
//}

func (m *defaultPlaceModel) GetPlaceList(in *pms.PlaceListReq) ([]*Place, int64, error) {
	var list []*Place
	db := m.conn.Model(&Place{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.PrincipalName != "" {
		db = db.Where("principal_name LIKE ?", fmt.Sprintf("%%%s%%", in.PrincipalName))
	}
	if in.PrincipalPhone != "" {
		db = db.Where("principal_phone LIKE ?", fmt.Sprintf("%%%s%%", in.PrincipalPhone))
	}
	if in.Address != "" {
		db = db.Where("address LIKE ?", fmt.Sprintf("%%%s%%", in.Address))
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

func (m *defaultPlaceModel) GetPlaceByID(id int64) (user *Place, err error) {
	err = m.conn.Model(&Place{}).Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
