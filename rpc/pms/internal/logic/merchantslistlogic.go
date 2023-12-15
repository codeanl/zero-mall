package logic

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsListLogic {
	return &MerchantsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商家列表
func (l *MerchantsListLogic) MerchantsList(in *pms.MerchantsListReq) (*pms.MerchantsListResp, error) {
	all, total, err := l.svcCtx.MerchantsModel.GetMerchantsList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.MerchantsListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		userInfo, _ := l.svcCtx.UserModel.GetUserByID(i.UserID)
		var info *pms.UserInfoFF
		jsonData1, _ := json.Marshal(userInfo)
		err = json.Unmarshal(jsonData1, &info)
		list[index].User = info
	}
	return &pms.MerchantsListResp{
		Total: total,
		List:  list,
	}, nil
}
