package place

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceInfoLogic {
	return &PlaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceInfoLogic) PlaceInfo(req *types.PlaceInfoReq) (resp *types.PlaceInfoResp, err error) {
	data, err := l.svcCtx.Pms.PlaceInfo(l.ctx, &pmsclient.PlaceInfoReq{
		Id: req.ID,
	})
	if err != nil {
		return &types.PlaceInfoResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var info *types.PlaceInfo
	jsonData, err := json.Marshal(data)
	err = json.Unmarshal(jsonData, &info)
	return &types.PlaceInfoResp{
		Code:    200,
		Message: "查询成功",
		Data:    info,
	}, nil
}
