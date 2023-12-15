package merchants

import (
	"context"
	"encoding/json"
	"fmt"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsInfoLogic {
	return &MerchantsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsInfoLogic) MerchantsInfo(req *types.MerchantsInfoReq) (resp *types.MerchantsInfoResp, err error) {
	data, err := l.svcCtx.Pms.MerchantsInfo(l.ctx, &pmsclient.MerchantsInfoReq{
		Id: req.ID,
	})
	fmt.Println(data.Id)
	if err != nil {
		return &types.MerchantsInfoResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var info *types.MerchantsInfo
	jsonData, err := json.Marshal(data)
	err = json.Unmarshal(jsonData, &info)
	return &types.MerchantsInfoResp{
		Code:    200,
		Message: "查询成功",
		Data:    info,
	}, nil
}
