package category

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
	"simple_mall_new/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateCategoryLogic {
	return &SaveOrUpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateCategoryLogic) SaveOrUpdateCategory(req *types.SaveOrUpdateCategoryReq) (resp *types.SaveOrUpdateCategoryResp, err error) {
	data := &pmsclient.SaveOrUpdateCategoryReq{
		Id:       req.ID,
		ParentId: req.ParentId,
		Name:     req.Name,
		Status:   req.Status,
		Sort:     req.Sort,
		Icon:     req.Icon,
		Desc:     req.Desc,
	}
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	//
	_, err = l.svcCtx.Pms.SaveOrUpdateCategory(l.ctx, data)
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateCategoryResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	//
	return &types.SaveOrUpdateCategoryResp{
		Code:    200,
		Message: msg,
	}, nil
}
