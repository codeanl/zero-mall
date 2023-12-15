package merchants_apply

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyListLogic {
	return &MerchantsApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsApplyListLogic) MerchantsApplyList(req *types.ListMerchantsApplyReq) (resp *types.ListMerchantsApplyResp, err error) {
	data, err := l.svcCtx.Pms.MerchantsApplyList(l.ctx, &pmsclient.MerchantsApplyListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
		Type:     req.Type,
	})
	if err != nil {
		return &types.ListMerchantsApplyResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListMerchantsApplyData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)
	return &types.ListMerchantsApplyResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
