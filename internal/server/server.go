package server

import (
	"github.com/zhoushuguang/rose/common/net/chttp"
	"github.com/zhoushuguang/rose/internal/conf"
	"github.com/zhoushuguang/rose/internal/handler"
	"github.com/zhoushuguang/rose/internal/service"
)

func NewHTTP(conf *conf.Conf) *chttp.Server {
	s := chttp.NewServer(conf.Server)
	svc := service.NewService(conf)

	handler.InitRouter(s, svc)

	err := s.Start()
	if err != nil {
		panic(err)
	}

	return s
}
