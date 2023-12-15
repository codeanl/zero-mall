package logic

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceInfoLogic {
	return &PlaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自提点详情
func (l *PlaceInfoLogic) PlaceInfo(in *pms.PlaceInfoReq) (*pms.PlaceInfoResp, error) {
	info, err := l.svcCtx.PlaceModel.GetPlaceByID(in.Id)
	if err != nil {
		return nil, err
	}
	var PlaceInfo *pms.PlaceInfoResp
	jsonData, _ := json.Marshal(info)
	_ = json.Unmarshal(jsonData, &PlaceInfo)
	MerchantsApplyInfo, err := l.svcCtx.MerchantsApplyModel.GetMerchantsApplyByID(info.MerchantApplysListID)
	//
	var MerchantsApplyInfoData *pms.MerchantApplysListData
	jsonData1, _ := json.Marshal(MerchantsApplyInfo)
	_ = json.Unmarshal(jsonData1, &MerchantsApplyInfoData)
	PlaceInfo.MerchantsApplyInfo = MerchantsApplyInfoData

	//
	return &pms.PlaceInfoResp{
		Id:                 PlaceInfo.Id,
		Name:               PlaceInfo.Name,
		PrincipalName:      PlaceInfo.PrincipalName,
		PrincipalPhone:     PlaceInfo.PrincipalPhone,
		Address:            PlaceInfo.Address,
		Pic:                PlaceInfo.Pic,
		MerchantsApplyInfo: PlaceInfo.MerchantsApplyInfo,
	}, nil
}
