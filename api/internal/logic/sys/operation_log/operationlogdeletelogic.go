package operation_log

import (
	"context"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperationLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationLogDeleteLogic {
	return &OperationLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperationLogDeleteLogic) OperationLogDelete(req *types.OperationLogDeleteReq) (resp *types.OperationLogDeleteResp, err error) {
	_, err = l.svcCtx.Sys.OperationLogDelete(l.ctx, &sysclient.OperationLogDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.OperationLogDeleteResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.OperationLogDeleteResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
