// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type SaveOrUpdateUserReq struct {
	ID       int64      `json:"id,optional"`
	Username string     `json:"username,optional"`
	Phone    string     `json:"phone,optional"`
	Nickname string     `json:"nickname,optional"`
	Gender   string     `json:"gender,optional"`
	Status   string     `json:"status,optional"`
	Email    string     `json:"email,optional"`
	Avatar   string     `json:"avatar,optional"`
	UserRole []UserRole `json:"user_role,optional"`
}

type UserRole struct {
	RoleID   int64  `json:"role_id"`
	DataType string `json:"data_type"`
}

type SaveOrUpdateUserResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteUserReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteUserResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListUserReq struct {
	PageNum  int64  `form:"page_num,default=1"`
	PageSize int64  `form:"page_size,default=10"`
	Nickname string `form:"nickname,optional"`
	Phone    string `form:"phone,optional"`
	Username string `form:"username,optional"`
	Status   string `form:"status,optional"`
	Gander   string `form:"gender,optional"`
	Email    string `form:"email,optional"`
}

type ListUser struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	NickName  string   `json:"nickname"`
	Phone     string   `json:"phone"`
	Gander    string   `json:"gender"`
	Avatar    string   `json:"avatar"`
	Email     string   `json:"email"`
	Status    string   `json:"status"`
	CreatedAt string   `json:"created_at"`
	RoleName  []string `json:"role_name"`
}

type ListUserResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    []*ListUser `json:"data"`
	Total   int64       `json:"total"`
}

type UpdatePasswordReq struct {
	ID          int64  `json:"id,optional"`
	Type        string `json:"type"`
	OldPassword string `json:"old_password,optional"`
	NewPassword string `json:"new_password,optional"`
}

type UpdatePasswordResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UserInfoResp struct {
	Code    int64        `json:"code"`
	Message string       `json:"message"`
	Data    UserInfoData `json:"data"`
}

type UserInfoData struct {
	User  Userr    `json:"user"`
	Roles []*Roles `json:"roles"`
}

type Userr struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
	Gander   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	CreateAt string `json:"create_at"`
}

type Roles struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type SaveOrUpdateRoleReq struct {
	ID     int64  `json:"id,optional"`
	Name   string `json:"name"`
	Remark string `json:"remark,optional"`
}

type SaveOrUpdateRoleResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteRoleReq struct {
	IDs []int64 `json:"ids"`
}

type DeleteRoleResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListRoleReq struct {
	PageNum  int64  `form:"page_num,optional"`
	PageSize int64  `form:"page_size,optional"`
	Name     string `form:"name,optional"`
}

type ListRoleData struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Remark   string `json:"remark"`
	CreateBy string `json:"create_by"`
	CreateAt string `json:"create_at"`
	UpdateBy string `json:"update_by"`
	UpdateAt string `json:"update_at"`
}

type ListRoleResp struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Total   int64           `json:"total"`
	Data    []*ListRoleData `json:"data"`
}

type SaveOrUpdateMenuReq struct {
	ID       int64  `json:"id,optional"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
	Url      string `json:"url"`
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"order_num"`
	TAG      string `json:"tag"`
}

type SaveOrUpdateMenuResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListMenuReq struct {
}

type ListMenuData struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	ParentId int64          `json:"parent_id"`
	Url      string         `json:"url"`
	Type     string         `json:"type"`
	Icon     string         `json:"icon"`
	OrderNum int64          `json:"order_num"`
	TAG      string         `json:"tag"`
	Children []ListMenuData `json:"children"`
}

type ListMenuResp struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Data    []ListMenuData `json:"data"`
}

type DeleteMenuReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteMenuResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListLoginLogReq struct {
	PageNum  int64 `form:"page_num,default=1"`
	PageSize int64 `form:"page_size,default=10"`
	UserID   int64 `form:"user_id,optional"`
}

type ListLoginLogData struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	IP        string `json:"ip"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}

type ListLoginLogResp struct {
	Code    int64               `json:"code"`
	Message string              `json:"message"`
	Total   int64               `json:"total"`
	Data    []*ListLoginLogData `json:"data"`
}

type DeleteLoginLogReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteLoginLogResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type OperationLogListReq struct {
	PageNum  int64  `form:"page_num,default=1"`
	PageSize int64  `form:"page_size,default=20"`
	Method   string `form:"method,optional"`
}

type ListSysLogData struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Operation string `json:"operation"`
	Method    string `json:"method"`
	Params    string `json:"params"`
	Time      int64  `json:"time"`
	IP        string `json:"ip"`
	CreatedAt string `json:"created_at"`
	UserInfo  User   `json:"user"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type OperationLogListResp struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Data    []*ListSysLogData `json:"data"`
	Total   int64             `json:"total"`
}

type OperationLogDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type OperationLogDeleteResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type SaveOrUpdateMemberReq struct {
	ID       int64  `json:"id,optional"`
	Password string `json:"password,optional"`
	Nickname string `json:"nickname,optional"`
	Phone    string `json:"phone,optional"`
	Status   string `json:"status,optional"`
	Avatar   string `json:"avatar,optional"`
	Gender   string `json:"gender,optional"`
	Email    string `json:"email,optional"`
}

type SaveOrUpdateMemberResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListMemberReq struct {
	PageNum  int64  `form:"page_num,default=1"`
	PageSize int64  `form:"page_size,default=20"`
	Nickname string `form:"nickname,optional"`
	Phone    string `form:"phone,optional"`
	Status   string `form:"status,optional"`
}

type ListMemberData struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	CreateAt string `json:"created_at"`
}

type ListMemberResp struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Data    []*ListMemberData `json:"data"`
	Total   int64             `json:"total"`
}

type DeleteMemberReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteMemberResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type SaveOrUpdateHomeAdvertiseReq struct {
	ID     int64  `json:"id,optional"`
	Name   string `json:"name,optional"`
	Pic    string `json:"pic,optional"`
	Status string `json:"status,optional"`
	Sort   int64  `json:"sort,optional"`
	Url    string `json:"url,optional"`
	Note   string `json:"note,optional"`
}

type SaveOrUpdateHomeAdvertiseResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListHomeAdvertiseReq struct {
	PageNum  int64  `form:"page_num,default=1"`
	PageSize int64  `form:"page_size,default=20"`
	Name     string `form:"name,optional"`
	Status   string `form:"status,optional"`
}

type ListHomeAdvertiseData struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	Status    string `json:"status"`
	Sort      int64  `json:"sort"`
	Url       string `json:"url"`
	Note      string `json:"note"`
	CreatedAt string `json:"created_at"`
}

type ListHomeAdvertiseResp struct {
	Code    int64                    `json:"code"`
	Message string                   `json:"message"`
	Data    []*ListHomeAdvertiseData `json:"data"`
	Total   int64                    `json:"total"`
}

type DeleteHomeAdvertiseReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteHomeAdvertiseResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type SaveOrUpdateCouponReq struct {
	ID          int64   `json:"id,optional"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Count       int64   `json:"count"`
	Amount      float64 `json:"amount"`
	PerLimit    int64   `json:"per_limit"`
	MinPoint    float64 `json:"min_point"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Note        string  `json:"note,optional"`
	EnableTime  string  `json:"enable_time"`
	Code        string  `json:"code"`
	UseType     string  `json:"use_type"`
	UseProduct  int64   `json:"use_product,optional"`
	UseCategory int64   `json:"use_category,optional"`
}

type SaveOrUpdateCouponResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListCouponReq struct {
	PageNum  int64  `form:"page_num,default=1"`
	PageSize int64  `form:"page_size,default=20"`
	Type     string `form:"type,optional"`
	Name     string `form:"name,optional"`
	UseType  string `form:"use_type,optional"`
}

type ListCouponData struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Count       int64   `json:"count"`
	Amount      float64 `json:"amount"`
	PerLimit    int64   `json:"per_limit"`
	MinPoint    float64 `json:"min_point"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Note        string  `json:"note,optional"`
	EnableTime  string  `json:"enable_time"`
	Code        string  `json:"code"`
	UseType     string  `json:"use_type"`
	UseProduct  int64   `json:"use_product"`
	UseCategory int64   `json:"use_category"`
	CreatedAt   string  `json:"created_at"`
}

type ListCouponResp struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Data    []*ListCouponData `json:"data"`
	Total   int64             `json:"total"`
}

type DeleteCouponReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteCouponResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type SaveOrUpdateSubjectReq struct {
	ID     int64  `json:"id,optional"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Status string `json:"status"`
	Sort   int64  `json:"sort"`
}

type SaveOrUpdateSubjectResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteSubjectReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteSubjectResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListSubjectReq struct {
	PageNum  int64  `form:"page_num,default=1"`
	PageSize int64  `form:"page_size,default=20"`
	Name     string `form:"name,optional "`
	Status   string `form:"status,optional"`
}

type ListSubjectData struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Status string `json:"status"`
	Sort   int64  `json:"sort"`
}

type ListSubjectResp struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Total   int64              `json:"total"`
	Data    []*ListSubjectData `json:"data"`
}

type ListCategoryReq struct {
	PageNum  int64  `json:"page_num,optional"`
	PageSize int64  `json:"page_size,optional"`
	Name     string `json:"name,optional"`
	ParentId int64  `json:"parent_id,optional"`
}

type ListCategoryData struct {
	ID       int64              `json:"id"`
	ParentId int64              `json:"parent_id"`
	Name     string             `json:"name"`
	Status   string             `json:"status"`
	Sort     int64              `json:"sort"`
	Icon     string             `json:"icon"`
	Desc     string             `json:"desc"`
	Children []ListCategoryData `json:"children"`
}

type ListCategoryResp struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Data    []ListCategoryData `json:"data"`
	Total   int64              `json:"total"`
}

type SaveOrUpdateCategoryReq struct {
	ID       int64  `json:"id,optional"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Sort     int64  `json:"sort"`
	Icon     string `json:"icon"`
	Desc     string `json:"desc,optional"`
}

type SaveOrUpdateCategoryResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteCategoryReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteCategoryResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type AddMerchantsApplyReq struct {
	PrincipalName  string `json:"principal_name"`
	PrincipalPhone string `json:"principal_phone"`
	IDCardFront    string `json:"id_card_front"`
	IDCardReverse  string `json:"id_card_reverse"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Type           string `json:"type"`
	Pic            string `json:"pic"`
	Remark         string `json:"remark,optional"`
}

type AddMerchantsApplyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListMerchantsApplyReq struct {
	PageNum  int64  `form:"page_num,default=1"`
	PageSize int64  `form:"page_size,default=20"`
	Name     string `form:"name,optional"`
	Status   string `form:"status,optional"`
	Type     string `form:"type,optional"`
}

type ListMerchantsApplyData struct {
	ID             int64  `json:"id"`
	PrincipalName  string `json:"principal_name"`
	PrincipalPhone string `json:"principal_phone"`
	IDCardFront    string `json:"id_card_front"`
	IDCardReverse  string `json:"id_card_reverse"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Pic            string `json:"pic"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	Auditor        string `json:"auditor"`
	ApprovalTime   string `json:"approval_time"`
	Remark         string `json:"remark"`
	AdminRemark    string `json:"admin_remark"`
	CreatedAt      string `json:"created_at"`
}

type ListMerchantsApplyResp struct {
	Code    int64                     `json:"code"`
	Message string                    `json:"message"`
	Data    []*ListMerchantsApplyData `json:"data"`
	Total   int64                     `json:"total"`
}

type UpdateMerchantsApplyReq struct {
	ID          int64  `json:"id"`
	Status      string `json:"status"`
	AdminRemark string `json:"admin_remark,optional"`
}

type UpdateMerchantsApplyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteMerchantsApplyReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteMerchantsApplyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListMerchantsReq struct {
	PageNum        int64  `form:"page_num,default=1"`
	PageSize       int64  `form:"page_size,default=20"`
	Name           string `form:"name,optional"`
	Address        string `form:"address,optional"`
	PrincipalPhone string `form:"principal_phone,optional"`
	PrincipalName  string `form:"principal_name,optional"`
}

type ListMerchantsData struct {
	ID             int64      `json:"id"`
	Name           string     `json:"name"`
	PrincipalPhone string     `json:"principal_phone"`
	PrincipalName  string     `json:"principal_name"`
	Address        string     `json:"address"`
	Pic            string     `json:"pic"`
	UserID         int64      `json:"user_id"`
	UserInfoFF     UserInfoFF `json:"user"`
}

type UserInfoFF struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type ListMerchantsResp struct {
	Code    int64                `json:"code"`
	Message string               `json:"message"`
	Data    []*ListMerchantsData `json:"data"`
	Total   int64                `json:"total"`
}

type UpdateMerchantsReq struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	PrincipalPhone string `json:"principal_phone"`
	PrincipalName  string `json:"principal_name"`
	Address        string `json:"address"`
	Pic            string `json:"pic"`
}

type UpdateMerchantsResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteMerchantsReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteMerchantsResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type MerchantsInfoReq struct {
	ID int64 `json:"id"`
}

type MerchantsInfoResp struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Data    *MerchantsInfo `json:"data"`
}

type MerchantsInfo struct {
	ID                 int64                  `json:"id"`
	Name               string                 `json:"name"`
	PrincipalPhone     string                 `json:"principal_phone"`
	PrincipalName      string                 `json:"principal_name"`
	Address            string                 `json:"address"`
	Pic                string                 `json:"pic"`
	UserID             int64                  `json:"user_id"`
	MerchantsApplyInfo MerchantApplysListData `json:"merchants_apply_info"`
}

type MerchantApplysListData struct {
	ID             int64  `json:"id"`
	PrincipalName  string `json:"principal_name"`
	PrincipalPhone string `json:"principal_phone"`
	IDCardFront    string `json:"id_card_front"`
	IDCardReverse  string `json:"id_card_reverse"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Pic            string `json:"pic"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	Auditor        string `json:"auditor"`
	ApprovalTime   string `json:"approval_time"`
	Remark         string `json:"remark"`
	AdminRemark    string `json:"admin_remark"`
	CreatedAt      string `json:"created_at"`
}
