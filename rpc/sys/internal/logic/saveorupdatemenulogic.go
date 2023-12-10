package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sys/model"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateMenuLogic {
	return &SaveOrUpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveOrUpdateMenu 添加｜｜更新菜单
func (l *SaveOrUpdateMenuLogic) SaveOrUpdateMenu(in *sys.SaveOrUpdateMenuReq) (*sys.SaveOrUpdateMenuResp, error) {
	menu, exist, _ := l.svcCtx.MenuModel.GetMenuByName(in.Name)
	if exist && int64(menu.ID) != in.Id {
		return nil, errors.New("菜单名已存在")
	}
	err := l.svcCtx.MenuModel.SaveOrUpdateMenu(in.Id, &model.Menu{
		Name:     in.Name,
		ParentID: in.ParentId,
		Url:      in.Url,
		Icon:     in.Icon,
		Type:     in.Type,
		OrderNum: in.OrderNum,
		TAG:      in.Tag,
	})
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}
	return &sys.SaveOrUpdateMenuResp{}, nil
}
