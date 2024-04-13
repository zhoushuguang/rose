package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/zhoushuguang/rose/internal/dto"
)

func (s *Service) UserDetail(ctx context.Context, uid int64) (*dto.UserDetailResp, error) {
	user, err := s.repo.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return &dto.UserDetailResp{
		User: &dto.User{
			UserId:     uid,
			UserName:   user.UserName,
			CreateTime: user.CreateTime.Unix(),
			UpdateTime: user.UpdateTime.Unix(),
		},
	}, nil
}

func (s *Service) UserCreate(ctx context.Context, userName string) (*dto.UserCreateResp, error) {
	id, err := s.repo.CreateUser(ctx, userName)
	if err != nil {
		return nil, err
	}
	return &dto.UserCreateResp{
		UserId: id,
	}, nil
}

func (s *Service) UserList(ctx context.Context, userIds string) (*dto.UserListResp, error) {
	strIds := strings.Split(userIds, ",")
	uIds := make([]string, 0, len(strIds))
	for _, strId := range strIds {
		_, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			return nil, err
		}
		uIds = append(uIds, strId)
	}

	us, err := s.repo.GetUserByIDs(ctx, uIds)
	if err != nil {
		return nil, err
	}
	users := make([]*dto.User, 0, len(us))
	for _, u := range us {
		users = append(users, &dto.User{
			UserId:     u.ID,
			UserName:   u.UserName,
			CreateTime: u.CreateTime.Unix(),
			UpdateTime: u.UpdateTime.Unix(),
		})
	}
	return &dto.UserListResp{
		Users: users,
	}, nil
}
