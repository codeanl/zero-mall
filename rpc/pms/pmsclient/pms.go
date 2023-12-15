// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package pmsclient

import (
	"context"

	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CategoryDeleteReq        = pms.CategoryDeleteReq
	CategoryDeleteResp       = pms.CategoryDeleteResp
	CategoryListData         = pms.CategoryListData
	CategoryListReq          = pms.CategoryListReq
	CategoryListResp         = pms.CategoryListResp
	MerchantApplysListData   = pms.MerchantApplysListData
	MerchantsApplyAddReq     = pms.MerchantsApplyAddReq
	MerchantsApplyAddResp    = pms.MerchantsApplyAddResp
	MerchantsApplyDeleteReq  = pms.MerchantsApplyDeleteReq
	MerchantsApplyDeleteResp = pms.MerchantsApplyDeleteResp
	MerchantsApplyListReq    = pms.MerchantsApplyListReq
	MerchantsApplyListResp   = pms.MerchantsApplyListResp
	MerchantsApplyUpdateReq  = pms.MerchantsApplyUpdateReq
	MerchantsApplyUpdateResp = pms.MerchantsApplyUpdateResp
	MerchantsDeleteReq       = pms.MerchantsDeleteReq
	MerchantsDeleteResp      = pms.MerchantsDeleteResp
	MerchantsListData        = pms.MerchantsListData
	MerchantsListReq         = pms.MerchantsListReq
	MerchantsListResp        = pms.MerchantsListResp
	MerchantsUpdateReq       = pms.MerchantsUpdateReq
	MerchantsUpdateResp      = pms.MerchantsUpdateResp
	SaveOrUpdateCategoryReq  = pms.SaveOrUpdateCategoryReq
	SaveOrUpdateCategoryResp = pms.SaveOrUpdateCategoryResp
	UserInfoFF               = pms.UserInfoFF

	Pms interface {
		// 添加||更新分类
		SaveOrUpdateCategory(ctx context.Context, in *SaveOrUpdateCategoryReq, opts ...grpc.CallOption) (*SaveOrUpdateCategoryResp, error)
		// 分类列表
		CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
		// 删除分类
		CategoryDelete(ctx context.Context, in *CategoryDeleteReq, opts ...grpc.CallOption) (*CategoryDeleteResp, error)
		// 入驻申请
		MerchantsApplyAdd(ctx context.Context, in *MerchantsApplyAddReq, opts ...grpc.CallOption) (*MerchantsApplyAddResp, error)
		// 入驻申请列表
		MerchantsApplyList(ctx context.Context, in *MerchantsApplyListReq, opts ...grpc.CallOption) (*MerchantsApplyListResp, error)
		// 审核入驻申请
		MerchantsApplyUpdate(ctx context.Context, in *MerchantsApplyUpdateReq, opts ...grpc.CallOption) (*MerchantsApplyUpdateResp, error)
		// 删除入驻申请
		MerchantsApplyDelete(ctx context.Context, in *MerchantsApplyDeleteReq, opts ...grpc.CallOption) (*MerchantsApplyDeleteResp, error)
		// 商家列表
		MerchantsList(ctx context.Context, in *MerchantsListReq, opts ...grpc.CallOption) (*MerchantsListResp, error)
		// 更新商家
		MerchantsUpdate(ctx context.Context, in *MerchantsUpdateReq, opts ...grpc.CallOption) (*MerchantsUpdateResp, error)
		// 删除商家
		MerchantsDelete(ctx context.Context, in *MerchantsDeleteReq, opts ...grpc.CallOption) (*MerchantsDeleteResp, error)
	}

	defaultPms struct {
		cli zrpc.Client
	}
)

func NewPms(cli zrpc.Client) Pms {
	return &defaultPms{
		cli: cli,
	}
}

// 添加||更新分类
func (m *defaultPms) SaveOrUpdateCategory(ctx context.Context, in *SaveOrUpdateCategoryReq, opts ...grpc.CallOption) (*SaveOrUpdateCategoryResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SaveOrUpdateCategory(ctx, in, opts...)
}

// 分类列表
func (m *defaultPms) CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryList(ctx, in, opts...)
}

// 删除分类
func (m *defaultPms) CategoryDelete(ctx context.Context, in *CategoryDeleteReq, opts ...grpc.CallOption) (*CategoryDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryDelete(ctx, in, opts...)
}

// 入驻申请
func (m *defaultPms) MerchantsApplyAdd(ctx context.Context, in *MerchantsApplyAddReq, opts ...grpc.CallOption) (*MerchantsApplyAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyAdd(ctx, in, opts...)
}

// 入驻申请列表
func (m *defaultPms) MerchantsApplyList(ctx context.Context, in *MerchantsApplyListReq, opts ...grpc.CallOption) (*MerchantsApplyListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyList(ctx, in, opts...)
}

// 审核入驻申请
func (m *defaultPms) MerchantsApplyUpdate(ctx context.Context, in *MerchantsApplyUpdateReq, opts ...grpc.CallOption) (*MerchantsApplyUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyUpdate(ctx, in, opts...)
}

// 删除入驻申请
func (m *defaultPms) MerchantsApplyDelete(ctx context.Context, in *MerchantsApplyDeleteReq, opts ...grpc.CallOption) (*MerchantsApplyDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyDelete(ctx, in, opts...)
}

// 商家列表
func (m *defaultPms) MerchantsList(ctx context.Context, in *MerchantsListReq, opts ...grpc.CallOption) (*MerchantsListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsList(ctx, in, opts...)
}

// 更新商家
func (m *defaultPms) MerchantsUpdate(ctx context.Context, in *MerchantsUpdateReq, opts ...grpc.CallOption) (*MerchantsUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsUpdate(ctx, in, opts...)
}

// 删除商家
func (m *defaultPms) MerchantsDelete(ctx context.Context, in *MerchantsDeleteReq, opts ...grpc.CallOption) (*MerchantsDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsDelete(ctx, in, opts...)
}
