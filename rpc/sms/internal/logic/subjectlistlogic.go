package logic

import (
	"context"
	"encoding/json"

	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListLogic {
	return &SubjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SubjectList 专题列表
func (l *SubjectListLogic) SubjectList(in *sms.SubjectListReq) (*sms.SubjectListResp, error) {
	all, total, err := l.svcCtx.SubjectModel.GetSubjectList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.SubjectListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreatedAt = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return &sms.SubjectListResp{
		Total: total,
		List:  list,
	}, nil
}
