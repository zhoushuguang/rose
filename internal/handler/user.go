package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhoushuguang/rose/internal/dto"
)

func userDetail(c *gin.Context) {
	r := new(dto.UserDetailReq)
	if err := c.Bind(r); err != nil {
		return
	}
	ret, err := svc.UserDetail(c.Request.Context(), r.UserId)
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
