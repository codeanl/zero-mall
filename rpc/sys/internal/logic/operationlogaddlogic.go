package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sys/model"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperationLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationLogAddLogic {
	return &OperationLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加日志
func (l *OperationLogAddLogic) OperationLogAdd(in *sys.OperationLogAddReq) (*sys.OperationLogAddResp, error) {
	err := l.svcCtx.OperationLogModel.AddOperationLog(&model.OperationLog{
		UserID:    in.UserId,
		Operation: in.Operation,
		Method:    in.Method,
		Params:    in.Params,
		Time:      in.Time,
		IP:        in.Ip,
	})
	if err != nil {
		return nil, errors.New("添加日志失败")
	}
	return &sys.OperationLogAddResp{}, nil
}
