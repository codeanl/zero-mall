package login_log

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogListLogic) LoginLogList(req *types.ListLoginLogReq) (resp *types.ListLoginLogResp, err error) {
	data, err := l.svcCtx.Sys.LoginLogList(l.ctx, &sysclient.LoginLogListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		UserId:   req.UserID,
	})
	if err != nil {
		return &types.ListLoginLogResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}

	var list []*types.ListLoginLogData

	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)

	return &types.ListLoginLogResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
