package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	cconf "github.com/zhoushuguang/rose/common/conf"
	"github.com/zhoushuguang/rose/internal/conf"
	"github.com/zhoushuguang/rose/internal/server"
)

var filePath = flag.String("conf", "etc/config.yaml", "the config path")

func main() {
	flag.Parse()

	c := new(conf.Conf)
	if err := cconf.Unmarshal(*filePath, c); err != nil {
		panic(err)
	}
	srv := server.NewHTTP(c)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_ = srv.Shutdown(context.Background())
			return
		default:
			return
		}
	}
}
