package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/ums/internal/svc"
	"simple_mall_new/rpc/ums/model"
	"simple_mall_new/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateMemberLogic {
	return &SaveOrUpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加｜｜更新会员
func (l *SaveOrUpdateMemberLogic) SaveOrUpdateMember(in *ums.SaveOrUpdateMemberReq) (*ums.SaveOrUpdateMemberResp, error) {
	//查询会员是否存在
	memberInfo, exist, _ := l.svcCtx.MemberModel.GetMemberByPhone(in.Phone)
	if exist && int64(memberInfo.ID) != in.Id {
		return nil, errors.New("账户已存在")
	}
	//todo 密码加密
	member := &model.Member{
		Phone:    in.Phone,
		Password: in.Password,
		Nickname: in.Nickname,
		Gender:   in.Gender,
		Email:    in.Email,
		Status:   in.Status,
		Avatar:   in.Avatar,
	}
	err := l.svcCtx.MemberModel.SaveOrUpdateMember(in.Id, member)
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}
	return &ums.SaveOrUpdateMemberResp{}, nil
}
