package subject

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sms/smsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListLogic {
	return &SubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectListLogic) SubjectList(req *types.ListSubjectReq) (resp *types.ListSubjectResp, err error) {
	data, err := l.svcCtx.Sms.SubjectList(l.ctx, &smsclient.SubjectListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})
	if err != nil {
		return &types.ListSubjectResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListSubjectData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)

	return &types.ListSubjectResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
