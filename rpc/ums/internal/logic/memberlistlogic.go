package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"simple_mall_new/rpc/ums/internal/svc"
	"simple_mall_new/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberListLogic {
	return &MemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 会员列表
func (l *MemberListLogic) MemberList(in *ums.MemberListReq) (*ums.MemberListResp, error) {
	all, total, err := l.svcCtx.MemberModel.GetMemberList(in)
	if err != nil {
		return nil, err
	}
	var list []*ums.MemberListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreatedAt = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	fmt.Println(list[0].CreatedAt)
	return &ums.MemberListResp{
		Total: total,
		List:  list,
	}, nil
}
