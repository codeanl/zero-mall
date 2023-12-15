// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"simple_mall_new/rpc/pms/internal/logic"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"
)

type PmsServer struct {
	svcCtx *svc.ServiceContext
	pms.UnimplementedPmsServer
}

func NewPmsServer(svcCtx *svc.ServiceContext) *PmsServer {
	return &PmsServer{
		svcCtx: svcCtx,
	}
}

// 添加||更新分类
func (s *PmsServer) SaveOrUpdateCategory(ctx context.Context, in *pms.SaveOrUpdateCategoryReq) (*pms.SaveOrUpdateCategoryResp, error) {
	l := logic.NewSaveOrUpdateCategoryLogic(ctx, s.svcCtx)
	return l.SaveOrUpdateCategory(in)
}

// 分类列表
func (s *PmsServer) CategoryList(ctx context.Context, in *pms.CategoryListReq) (*pms.CategoryListResp, error) {
	l := logic.NewCategoryListLogic(ctx, s.svcCtx)
	return l.CategoryList(in)
}

// 删除分类
func (s *PmsServer) CategoryDelete(ctx context.Context, in *pms.CategoryDeleteReq) (*pms.CategoryDeleteResp, error) {
	l := logic.NewCategoryDeleteLogic(ctx, s.svcCtx)
	return l.CategoryDelete(in)
}

// 入驻申请
func (s *PmsServer) MerchantsApplyAdd(ctx context.Context, in *pms.MerchantsApplyAddReq) (*pms.MerchantsApplyAddResp, error) {
	l := logic.NewMerchantsApplyAddLogic(ctx, s.svcCtx)
	return l.MerchantsApplyAdd(in)
}

// 入驻申请列表
func (s *PmsServer) MerchantsApplyList(ctx context.Context, in *pms.MerchantsApplyListReq) (*pms.MerchantsApplyListResp, error) {
	l := logic.NewMerchantsApplyListLogic(ctx, s.svcCtx)
	return l.MerchantsApplyList(in)
}

// 审核入驻申请
func (s *PmsServer) MerchantsApplyUpdate(ctx context.Context, in *pms.MerchantsApplyUpdateReq) (*pms.MerchantsApplyUpdateResp, error) {
	l := logic.NewMerchantsApplyUpdateLogic(ctx, s.svcCtx)
	return l.MerchantsApplyUpdate(in)
}

// 删除入驻申请
func (s *PmsServer) MerchantsApplyDelete(ctx context.Context, in *pms.MerchantsApplyDeleteReq) (*pms.MerchantsApplyDeleteResp, error) {
	l := logic.NewMerchantsApplyDeleteLogic(ctx, s.svcCtx)
	return l.MerchantsApplyDelete(in)
}

// 商家列表
func (s *PmsServer) MerchantsList(ctx context.Context, in *pms.MerchantsListReq) (*pms.MerchantsListResp, error) {
	l := logic.NewMerchantsListLogic(ctx, s.svcCtx)
	return l.MerchantsList(in)
}

// 更新商家
func (s *PmsServer) MerchantsUpdate(ctx context.Context, in *pms.MerchantsUpdateReq) (*pms.MerchantsUpdateResp, error) {
	l := logic.NewMerchantsUpdateLogic(ctx, s.svcCtx)
	return l.MerchantsUpdate(in)
}

// 删除商家
func (s *PmsServer) MerchantsDelete(ctx context.Context, in *pms.MerchantsDeleteReq) (*pms.MerchantsDeleteResp, error) {
	l := logic.NewMerchantsDeleteLogic(ctx, s.svcCtx)
	return l.MerchantsDelete(in)
}

// 商家详情
func (s *PmsServer) MerchantsInfo(ctx context.Context, in *pms.MerchantsInfoReq) (*pms.MerchantsInfoResp, error) {
	l := logic.NewMerchantsInfoLogic(ctx, s.svcCtx)
	return l.MerchantsInfo(in)
}

// 自提点列表
func (s *PmsServer) PlaceList(ctx context.Context, in *pms.PlaceListReq) (*pms.PlaceListResp, error) {
	l := logic.NewPlaceListLogic(ctx, s.svcCtx)
	return l.PlaceList(in)
}

// 更新自提点
func (s *PmsServer) PlaceUpdate(ctx context.Context, in *pms.PlaceUpdateReq) (*pms.PlaceUpdateResp, error) {
	l := logic.NewPlaceUpdateLogic(ctx, s.svcCtx)
	return l.PlaceUpdate(in)
}

// 删除自提点
func (s *PmsServer) PlaceDelete(ctx context.Context, in *pms.PlaceDeleteReq) (*pms.PlaceDeleteResp, error) {
	l := logic.NewPlaceDeleteLogic(ctx, s.svcCtx)
	return l.PlaceDelete(in)
}

// 自提点详情
func (s *PmsServer) PlaceInfo(ctx context.Context, in *pms.PlaceInfoReq) (*pms.PlaceInfoResp, error) {
	l := logic.NewPlaceInfoLogic(ctx, s.svcCtx)
	return l.PlaceInfo(in)
}
