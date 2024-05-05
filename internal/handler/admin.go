package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhoushuguang/rose/internal/dto"
)

func adminPurcharseList(c *gin.Context) {
	r := new(dto.AdminPurcharseListReq)
	if err := c.Bind(r); err != nil {
		return
	}
	if r.PageNo == 0 {
		r.PageNo = 1
	}
	if r.PageSize == 0 {
		r.PageSize = 10
	}
	ret, err := svc.AdminPurcharseList(c.Request.Context(), r.UserId, r.PageSize, r.PageNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, ret)
}
