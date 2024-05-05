package dto

type (
	AdminPurcharseListReq struct {
		UserId   int64 `form:"userId"`
		PageSize int   `form:"pageSize"`
		PageNo   int   `form:"pageNo"`
	}

	Purcharse struct {
		UserId     int64   `json:"userId"`
		UserName   string  `json:"userName"`
		Price      float64 `json:"price"`
		CreateTime int64   `json:"createTime"`
	}

	AdminPurcharseListResp struct {
		Purcharses []*Purcharse `json:"purcharses"`
		Count      int          `json:"count"`
		PageNo     int          `json:"pageNo"`
	}
)
