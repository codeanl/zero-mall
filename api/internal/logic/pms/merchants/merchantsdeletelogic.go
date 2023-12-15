package merchants

import (
	"context"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsDeleteLogic {
	return &MerchantsDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsDeleteLogic) MerchantsDelete(req *types.DeleteMerchantsReq) (resp *types.DeleteMerchantsResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsDelete(l.ctx, &pmsclient.MerchantsDeleteReq{
		IDs: req.Ids,
	})
	if err != nil {
		return &types.DeleteMerchantsResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteMerchantsResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
