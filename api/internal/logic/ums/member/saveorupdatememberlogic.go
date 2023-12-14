package member

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
	"simple_mall_new/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateMemberLogic {
	return &SaveOrUpdateMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateMemberLogic) SaveOrUpdateMember(req *types.SaveOrUpdateMemberReq) (resp *types.SaveOrUpdateMemberResp, err error) {
	_, err = l.svcCtx.Ums.SaveOrUpdateMember(l.ctx, &umsclient.SaveOrUpdateMemberReq{
		Id:       req.ID,
		Password: req.Password,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Status:   req.Status,
		Avatar:   req.Avatar,
		Gender:   req.Gender,
		Email:    req.Email,
	})
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateMemberResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.SaveOrUpdateMemberResp{
		Code:    200,
		Message: msg,
	}, nil
}
