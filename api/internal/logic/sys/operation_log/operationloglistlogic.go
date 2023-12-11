package operation_log

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationLogListLogic {
	return &OperationLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperationLogListLogic) OperationLogList(req *types.OperationLogListReq) (resp *types.OperationLogListResp, err error) {
	data, err := l.svcCtx.Sys.OperationLogList(l.ctx, &sysclient.OperationLogListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Method:   req.Method,
	})
	if err != nil {
		return &types.OperationLogListResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListSysLogData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)

	return &types.OperationLogListResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
