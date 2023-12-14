package member

import (
	"context"
	"encoding/json"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
	"simple_mall_new/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberListLogic {
	return &MemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberListLogic) MemberList(req *types.ListMemberReq) (resp *types.ListMemberResp, err error) {
	data, err := l.svcCtx.Ums.MemberList(l.ctx, &umsclient.MemberListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Status:   req.Status,
	})
	if err != nil {
		return &types.ListMemberResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListMemberData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)

	return &types.ListMemberResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
