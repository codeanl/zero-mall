package user

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"
	"simple_mall_new/rpc/sys/sys"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateUserLogic {
	return &SaveOrUpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateUserLogic) SaveOrUpdateUser(req *types.SaveOrUpdateUserReq) (resp *types.SaveOrUpdateUserResp, err error) {
	var list []*sys.UserRole
	jsonData, _ := json.Marshal(req.UserRole)
	json.Unmarshal(jsonData, &list)
	//
	_, err = l.svcCtx.Sys.SaveOrUpdateUser(l.ctx, &sysclient.SaveOrUpdateUserReq{
		Id:       req.ID,
		Username: req.Username,
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Gender:   req.Gender,
		Email:    req.Email,
		Status:   req.Status,
		Avatar:   req.Avatar,
		UserRole: list,
	})
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateUserResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.SaveOrUpdateUserResp{
		Code:    200,
		Message: msg,
	}, nil
}
