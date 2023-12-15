package merchants_apply

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
	"simple_mall_new/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyUpdateLogic {
	return &MerchantsApplyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsApplyUpdateLogic) MerchantsApplyUpdate(req *types.UpdateMerchantsApplyReq) (resp *types.UpdateMerchantsApplyResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsApplyUpdate(l.ctx, &pmsclient.MerchantsApplyUpdateReq{
		Id:      req.ID,
		Status:  req.Status,
		Auditor: l.ctx.Value("username").(string),
		//ApprovalTime: time.Now().Format("2006-01-02 15:04:05"),
		AdminRemark: req.AdminRemark,
	})
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.UpdateMerchantsApplyResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.UpdateMerchantsApplyResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
