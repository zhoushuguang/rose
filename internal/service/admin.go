package service

import (
	"context"
	"strconv"

	"github.com/zhoushuguang/rose/internal/dto"
)

func (s *Service) AdminPurcharseList(ctx context.Context, userId int64, pageSize, pageNo int) (*dto.AdminPurcharseListResp, error) {
	ps, count, err := s.repo.GetPurchaseWithPage(ctx, userId, pageSize, pageNo)
	if err != nil {
		return nil, err
	}
	var userIds []string
	um := make(map[int64]struct{})
	for _, p := range ps {
		if _, ok := um[p.UserId]; ok {
			continue
		}
		um[p.UserId] = struct{}{}
		userIdStr := strconv.FormatInt(p.UserId, 10)
		userIds = append(userIds, userIdStr)
	}
	us, err := s.repo.GetUserByIDs(ctx, userIds)
	if err != nil {
		return nil, err
	}
	users := make(map[int64]string)
	for _, u := range us {
		users[u.ID] = u.UserName
	}
	ret := &dto.AdminPurcharseListResp{
		Count:      count,
		PageNo:     pageNo,
		Purcharses: make([]*dto.Purcharse, 0, len(ps)),
	}
	for _, p := range ps {
		ret.Purcharses = append(ret.Purcharses, &dto.Purcharse{
			UserId:     p.UserId,
			UserName:   users[p.UserId],
			Price:      p.Price,
			CreateTime: p.CreateTime.Unix(),
		})
	}

	return ret, nil
}
