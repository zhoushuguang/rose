package service

import (
	"context"

	"github.com/zhoushuguang/rose/internal/dto"
)

func (s *Service) UserDetail(ctx context.Context, uid int64) (*dto.UserDetailResp, error) {
	// ignore biz logic
	return &dto.UserDetailResp{
		UserId:   uid,
		UserName: "hello rose",
	}, nil
}
