package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseListLogic {
	return &HomeAdvertiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 广告列表
func (l *HomeAdvertiseListLogic) HomeAdvertiseList(in *sms.HomeAdvertiseListReq) (*sms.HomeAdvertiseListResp, error) {
	all, total, err := l.svcCtx.HomeAdvertiseModel.GetHomeAdvertiseList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.HomeAdvertiseListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreatedAt = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	fmt.Println(list[0].CreatedAt)
	return &sms.HomeAdvertiseListResp{
		Total: total,
		List:  list,
	}, nil
}
