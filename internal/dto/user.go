package dto

type UserDetailReq struct {
	Uid int64 `form:"uid"`
}

type UserDetailResp struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
}
