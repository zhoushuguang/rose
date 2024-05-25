package chttp

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
)

const userIDHeader = "X-User-ID"

type headerKey struct{}

// Header head信息.
type Header struct {
	UserID int64 `json:"userID"`
}

// HeaderHandler header信息写入context.
func HeaderHandler(ginCtx *gin.Context) {
	h := &Header{}

	userIDStr := ginCtx.GetHeader(userIDHeader)
	h.UserID, _ = strconv.ParseInt(userIDStr, 10, 64)

	ctx := ginCtx.Request.Context()
	ctx = context.WithValue(ctx, headerKey{}, h)

	ginCtx.Request = ginCtx.Request.WithContext(ctx)
	ginCtx.Next()
}

// GetHeader 获取head信息.
func GetHeader(ctx context.Context) *Header {
	h, _ := ctx.Value(headerKey{}).(*Header)
	return h
}

// GetUserID 获取用户ID.
func GetUserID(ctx context.Context) int64 {
	h := GetHeader(ctx)
	if h == nil {
		return 0
	}

	return h.UserID
}
