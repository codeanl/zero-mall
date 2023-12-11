package logic

import (
	"context"
	"encoding/json"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationLogListLogic {
	return &OperationLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 日志列表
func (l *OperationLogListLogic) OperationLogList(in *sys.OperationLogListReq) (*sys.OperationLogListResp, error) {
	all, total, err := l.svcCtx.OperationLogModel.GetOperationLogList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.SysLogListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, log := range list {
		userinfo, _ := l.svcCtx.UserModel.GetUserByID(log.UserId)
		var user *sys.User
		userInfo, _ := json.Marshal(userinfo)
		err = json.Unmarshal(userInfo, &user)
		list[index].User = user
	}

	return &sys.OperationLogListResp{
		Total: total,
		List:  list,
	}, nil
}
