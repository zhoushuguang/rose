package service

import (
	"github.com/zhoushuguang/rose/internal/conf"
	"github.com/zhoushuguang/rose/internal/repo"
)

type Service struct {
	conf *conf.Conf
	repo *repo.Repo
}

func NewService(conf *conf.Conf) *Service {
	return &Service{
		conf: conf,
		repo: repo.NewRepo(conf),
	}
}
