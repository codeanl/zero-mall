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
