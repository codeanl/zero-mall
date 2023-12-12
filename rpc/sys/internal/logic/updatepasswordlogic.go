package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/model"
	"simple_mall_new/rpc/sys/sys"
	"simple_mall_new/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新密码｜｜重置密码
func (l *UpdatePasswordLogic) UpdatePassword(in *sys.UpdatePasswordReq) (*sys.UpdatePasswordResp, error) {
	if in.Type == "0" { //更新密码
		if in.OldPassword != in.NewPassword {
			return nil, errors.New("两次密码不一致")
		}
		//获取信息得到密码
		my, _ := l.svcCtx.UserModel.GetUserByID(in.Id)
		yes := utils.CheckPassword(my.Password, in.OldPassword)
		if !yes {
			return nil, errors.New("旧密码不正确")
		}
		//更新
		err := l.svcCtx.UserModel.SaveOrUpdateUser(in.Id, &model.User{
			Password: utils.SetPassword(in.NewPassword),
		})
		if err != nil {
			return nil, errors.New("更新失败")
		}
	} else if in.Type == "1" { //重置密码
		//更新
		err := l.svcCtx.UserModel.SaveOrUpdateUser(in.Id, &model.User{
			Password: utils.SetPassword("123456"),
		})
		if err != nil {
			return nil, errors.New("重置失败")
		}
	}

	return &sys.UpdatePasswordResp{}, nil

}
