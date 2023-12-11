package logic

import (
	"context"
	"encoding/json"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginLogList 登录日志列表
func (l *LoginLogListLogic) LoginLogList(in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	all, total, err := l.svcCtx.LoginLogModel.GetLoginLogList(in)
	if err != nil {
		return nil, err
	}
	//
	var list []*sys.LoginLogListData
	jsonData, err := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	//
	return &sys.LoginLogListResp{
		Total: total,
		List:  list,
	}, nil

}
