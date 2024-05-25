package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhoushuguang/rose/common/net/chttp"
	"github.com/zhoushuguang/rose/internal/dto"
)

func userDetail(c *gin.Context) {
	userID := chttp.GetUserID(c.Request.Context())
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	ret, err := svc.UserDetail(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, ret)
}

func userCreate(c *gin.Context) {
	r := new(dto.UserCreateReq)
	if err := c.BindJSON(r); err != nil {
		return
	}
	ret, err := svc.UserCreate(c.Request.Context(), r.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, ret)
}

func userList(c *gin.Context) {
	r := new(dto.UserListReq)
	if err := c.Bind(r); err != nil {
		return
	}
	ret, err := svc.UserList(c.Request.Context(), r.UserIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, ret)
}
