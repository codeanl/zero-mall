package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsInfoLogic {
	return &MerchantsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商家详情
func (l *MerchantsInfoLogic) MerchantsInfo(in *pms.MerchantsInfoReq) (*pms.MerchantsInfoResp, error) {
	info, err := l.svcCtx.MerchantsModel.GetMerchantsByID(in.Id)
	if err != nil {
		return nil, err
	}
	var MerchantsInfo *pms.MerchantsInfoResp
	jsonData, _ := json.Marshal(info)
	_ = json.Unmarshal(jsonData, &MerchantsInfo)
	MerchantsApplyInfo, err := l.svcCtx.MerchantsApplyModel.GetMerchantsApplyByID(info.MerchantApplysListID)
	//
	var MerchantsApplyInfoData *pms.MerchantApplysListData
	jsonData1, _ := json.Marshal(MerchantsApplyInfo)
	_ = json.Unmarshal(jsonData1, &MerchantsApplyInfoData)
	MerchantsInfo.MerchantsApplyInfo = MerchantsApplyInfoData
	fmt.Println(MerchantsInfo.MerchantsApplyInfo)

	//
	return &pms.MerchantsInfoResp{
		Id:                 MerchantsInfo.Id,
		Name:               MerchantsInfo.Name,
		PrincipalName:      MerchantsInfo.PrincipalName,
		PrincipalPhone:     MerchantsInfo.PrincipalPhone,
		Address:            MerchantsInfo.Address,
		Pic:                MerchantsInfo.Pic,
		MerchantsApplyInfo: MerchantsInfo.MerchantsApplyInfo,
	}, nil
}
