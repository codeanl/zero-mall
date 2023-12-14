package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/model"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateHomeAdvertiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateHomeAdvertiseLogic {
	return &SaveOrUpdateHomeAdvertiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加｜｜更新广告
func (l *SaveOrUpdateHomeAdvertiseLogic) SaveOrUpdateHomeAdvertise(in *sms.SaveOrUpdateHomeAdvertiseReq) (*sms.SaveOrUpdateHomeAdvertiseResp, error) {
	data := &model.HomeAdvertise{
		Name:   in.Name,
		Pic:    in.Pic,
		Status: in.Status,
		Sort:   in.Sort,
		Url:    in.Url,
		Note:   in.Note,
	}
	err := l.svcCtx.HomeAdvertiseModel.SaveOrUpdateHomeAdvertise(in.Id, data)
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}
	return &sms.SaveOrUpdateHomeAdvertiseResp{}, nil
}
