package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/model"
	"simple_mall_new/rpc/sms/sms"
	"time"
)

type SaveOrUpdateCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateCouponLogic {
	return &SaveOrUpdateCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加｜｜更新优惠券
func (l *SaveOrUpdateCouponLogic) SaveOrUpdateCoupon(in *sms.SaveOrUpdateCouponReq) (*sms.SaveOrUpdateCouponResp, error) {
	//查询用户名是否存在
	userInfo, exist, _ := l.svcCtx.CouponModel.GetCouponByName(in.Name)
	if exist && int64(userInfo.ID) != in.Id {
		return nil, errors.New("优惠券已存在")
	}
	//
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	EnableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)
	coupon := &model.Coupon{
		Type:        in.Type,
		Name:        in.Name,
		Count:       in.Count,
		Amount:      float64(in.Amount),
		PerLimit:    in.PerLimit,
		MinPoint:    float64(in.MinPoint),
		StartTime:   StartTime,
		EndTime:     endTime,
		Note:        in.Note,
		EnableTime:  EnableTime,
		Code:        in.Code,
		UseType:     in.UseType,
		UseProduct:  in.UseProduct,
		UseCategory: in.UseCategory,
	}

	// 将字符串转换为time.Time类型
	err := l.svcCtx.CouponModel.SaveOrUpdateCoupon(in.Id, coupon)
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}

	return &sms.SaveOrUpdateCouponResp{}, nil
}
