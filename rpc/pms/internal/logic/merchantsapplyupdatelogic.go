package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/model"
	"simple_mall_new/rpc/pms/pms"
	sysModel "simple_mall_new/rpc/sys/model"
	"simple_mall_new/utils"
	"time"
)

type MerchantsApplyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyUpdateLogic {
	return &MerchantsApplyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 审核入驻申请
func (l *MerchantsApplyUpdateLogic) MerchantsApplyUpdate(in *pms.MerchantsApplyUpdateReq) (*pms.MerchantsApplyUpdateResp, error) {
	//获取信息
	info, _ := l.svcCtx.MerchantsApplyModel.GetMerchantsApplyByID(in.Id)
	ApprovalTime, _ := time.Parse("2006-01-02 15:04:05", in.ApprovalTime)
	err := l.svcCtx.MerchantsApplyModel.UpdateMerchantsApply(in.Id, &model.MerchantsApply{
		Status:       in.Status,
		Auditor:      in.Auditor,
		ApprovalTime: ApprovalTime,
		AdminRemark:  in.AdminRemark,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}

	//
	if info.Status == "1" || info.Status == "2" {
		return nil, errors.New("该入驻申请已经处理，无需处理")
	}
	//如果审核通过
	if in.Status == "1" {
		//创建用户
		userInfo, err := l.svcCtx.UserModel.SaveOrUpdateUser(0, &sysModel.User{
			Username: info.PrincipalPhone,
			Phone:    info.PrincipalPhone,
			Nickname: info.PrincipalName,
			Gender:   "0",
			Avatar:   info.Pic,
			Email:    "0",
			Status:   "1",
			Password: utils.SetPassword("123456"),
		})
		if err != nil {
			return nil, errors.New("入驻创建用户失败")
		}
		fmt.Println(userInfo.ID)

		//0->商家 1->自提点
		if info.Type == "0" {
			_, err := l.svcCtx.MerchantsModel.AddMerchants(&model.Merchants{
				UserID:               int64(userInfo.ID),
				Name:                 info.Name,
				PrincipalName:        info.PrincipalName,
				PrincipalPhone:       info.PrincipalPhone,
				Address:              info.Address,
				Pic:                  info.Pic,
				MerchantApplysListID: int64(info.ID),
			})
			if err != nil {
				return nil, errors.New("入驻失败")
			}
		} else if info.Type == "1" {
			err := l.svcCtx.PlaceModel.AddPlace(&model.Place{
				UserID:               int64(userInfo.ID),
				Name:                 info.Name,
				PrincipalName:        info.PrincipalName,
				PrincipalPhone:       info.PrincipalPhone,
				Address:              info.Address,
				Pic:                  info.Pic,
				MerchantApplysListID: int64(info.ID),
			})
			if err != nil {
				return nil, errors.New("入驻失败")
			}
		}
	}
	return &pms.MerchantsApplyUpdateResp{}, nil
}
