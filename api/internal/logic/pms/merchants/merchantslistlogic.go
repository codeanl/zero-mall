package merchants

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsListLogic {
	return &MerchantsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsListLogic) MerchantsList(req *types.ListMerchantsReq) (resp *types.ListMerchantsResp, err error) {
	data, err := l.svcCtx.Pms.MerchantsList(l.ctx, &pmsclient.MerchantsListReq{
		PageNum:        req.PageNum,
		PageSize:       req.PageSize,
		Name:           req.Name,
		Address:        req.Address,
		PrincipalPhone: req.PrincipalPhone,
		PrincipalName:  req.PrincipalName,
	})
	if err != nil {
		return &types.ListMerchantsResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListMerchantsData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)
	return &types.ListMerchantsResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
