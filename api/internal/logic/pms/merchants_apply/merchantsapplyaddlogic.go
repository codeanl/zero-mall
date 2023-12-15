package merchants_apply

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsApplyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyAddLogic {
	return &MerchantsApplyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsApplyAddLogic) MerchantsApplyAdd(req *types.AddMerchantsApplyReq) (resp *types.AddMerchantsApplyResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsApplyAdd(l.ctx, &pmsclient.MerchantsApplyAddReq{
		PrincipalName:  req.PrincipalName,
		PrincipalPhone: req.PrincipalPhone,
		IdCardFront:    req.IDCardFront,
		IdCardReverse:  req.IDCardReverse,
		Name:           req.Name,
		Address:        req.Address,
		Pic:            req.Pic,
		Type:           req.Type,
		Status:         "0",
		Remark:         req.Remark,
	})
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.AddMerchantsApplyResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.AddMerchantsApplyResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
