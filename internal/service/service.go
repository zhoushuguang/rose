package service

import (
	"github.com/zhoushuguang/rose/internal/conf"
)

type Service struct {
	conf *conf.Conf
}

func NewService(conf *conf.Conf) *Service {
	return &Service{
		conf: conf,
	}
}
