package logic

import (
	"context"
	"errors"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperationLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationLogDeleteLogic {
	return &OperationLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除日志
func (l *OperationLogDeleteLogic) OperationLogDelete(in *sys.OperationLogDeleteReq) (*sys.OperationLogDeleteResp, error) {
	err := l.svcCtx.OperationLogModel.DeleteOperationLogByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &sys.OperationLogDeleteResp{}, nil
}
