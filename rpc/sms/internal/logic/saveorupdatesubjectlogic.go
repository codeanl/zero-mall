package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/model"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateSubjectLogic {
	return &SaveOrUpdateSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveOrUpdateSubject 添加｜｜更新专题
func (l *SaveOrUpdateSubjectLogic) SaveOrUpdateSubject(in *sms.SaveOrUpdateSubjectReq) (*sms.SaveOrUpdateSubjectResp, error) {
	//查询用户名是否存在
	userInfo, exist, _ := l.svcCtx.SubjectModel.GetSubjectByName(in.Name)
	if exist && int64(userInfo.ID) != in.Id {
		return nil, errors.New("专题已存在")
	}
	err := l.svcCtx.SubjectModel.SaveOrUpdateCoupon(in.Id, &model.Subject{
		Name:   in.Name,
		Pic:    in.Pic,
		Status: in.Status,
		Sort:   in.Sort,
	})
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}

	return &sms.SaveOrUpdateSubjectResp{}, nil
}
