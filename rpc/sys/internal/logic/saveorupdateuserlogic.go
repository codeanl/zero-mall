package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sys/model"
	"simple_mall_new/utils"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateUserLogic {
	return &SaveOrUpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveOrUpdateUser 添加｜｜更新用户
func (l *SaveOrUpdateUserLogic) SaveOrUpdateUser(in *sys.SaveOrUpdateUserReq) (*sys.SaveOrUpdateUserResp, error) {
	//查询用户名是否存在
	userInfo, exist, _ := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if exist && int64(userInfo.ID) != in.Id {
		return nil, errors.New("账户已存在")
	}
	//
	user := &model.User{
		Username: in.Username,
		Phone:    in.Phone,
		Nickname: in.Nickname,
		Gender:   in.Gender,
		Email:    in.Email,
		Status:   in.Status,
		Avatar:   in.Avatar,
	}
	if in.Id == 0 {
		user.Password = utils.SetPassword("123456")
	}
	err := l.svcCtx.UserModel.SaveOrUpdateUser(in.Id, user)
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}
	//添加用户角色

	if len(in.UserRole) > 0 && in.Id > 0 {
		err = l.svcCtx.UserRoleModel.DeleteByUserID(in.Id)
		if err != nil {
			return nil, errors.New("系统错误")
		}
		for _, i := range in.UserRole {
			err = l.svcCtx.UserRoleModel.AddUserRole(&model.UserRole{
				UserID:   in.Id,
				RoleID:   i.RoleId,
				DataType: i.DataType,
			})
			if err != nil {
				return nil, errors.New("添加用户角色失败")
			}
		}
	}
	return &sys.SaveOrUpdateUserResp{}, nil
}
