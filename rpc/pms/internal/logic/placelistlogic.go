package logic

import (
	"context"
	"encoding/json"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceListLogic {
	return &PlaceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自提点列表
func (l *PlaceListLogic) PlaceList(in *pms.PlaceListReq) (*pms.PlaceListResp, error) {
	all, total, err := l.svcCtx.PlaceModel.GetPlaceList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.PlaceListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		userInfo, _ := l.svcCtx.UserModel.GetUserByID(i.UserID)
		var info *pms.UserInfoFF
		jsonData1, _ := json.Marshal(userInfo)
		err = json.Unmarshal(jsonData1, &info)
		list[index].User = info
	}
	return &pms.PlaceListResp{
		Total: total,
		List:  list,
	}, nil
}
