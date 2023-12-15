package place

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceListLogic {
	return &PlaceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceListLogic) PlaceList(req *types.ListPlaceReq) (resp *types.ListPlaceResp, err error) {
	data, err := l.svcCtx.Pms.PlaceList(l.ctx, &pmsclient.PlaceListReq{
		PageNum:        req.PageNum,
		PageSize:       req.PageSize,
		Name:           req.Name,
		Address:        req.Address,
		PrincipalPhone: req.PrincipalPhone,
		PrincipalName:  req.PrincipalName,
	})
	if err != nil {
		return &types.ListPlaceResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListPlaceData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)
	return &types.ListPlaceResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
