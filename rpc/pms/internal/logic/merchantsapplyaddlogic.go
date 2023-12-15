package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/pms/model"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyAddLogic {
	return &MerchantsApplyAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 入驻申请
func (l *MerchantsApplyAddLogic) MerchantsApplyAdd(in *pms.MerchantsApplyAddReq) (*pms.MerchantsApplyAddResp, error) {
	//查询是否存在
	_, exist, _ := l.svcCtx.MerchantsApplyModel.GetMerchantsApplyByPhone(in.PrincipalPhone)
	if exist {
		return nil, errors.New("号码已存在")
	}
	info := &model.MerchantsApply{
		PrincipalName:  in.PrincipalName,
		PrincipalPhone: in.PrincipalPhone,
		IDCardFront:    in.IdCardFront,
		IDCardReverse:  in.IdCardReverse,
		Name:           in.Name,
		Address:        in.Address,
		Pic:            in.Pic,
		Type:           in.Type,
		Status:         in.Status,
		Remark:         in.Remark,
	}
	_, err := l.svcCtx.MerchantsApplyModel.AddMerchantsApply(info)
	if err != nil {
		return nil, errors.New("入驻申请失败")
	}
	return &pms.MerchantsApplyAddResp{}, nil
}
