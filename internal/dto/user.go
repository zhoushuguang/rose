package dto

type (
	UserDetailReq struct {
		UserId int64 `form:"userId"`
	}

	UserDetailResp struct {
		*User
	}

	User struct {
		UserId     int64  `json:"userId"`
		UserName   string `json:"userName"`
		CreateTime int64  `json:"createTime"`
		UpdateTime int64  `json:"updateTime"`
	}

	UserCreateReq struct {
		UserName string `json:"userName"`
	}

	UserCreateResp struct {
		UserId int64 `json:"userId"`
	}

	UserListReq struct {
		UserIds string `form:"userIds"`
	}

	UserListResp struct {
		Users []*User `json:"users"`
	}
)
