package conf

import (
	"github.com/zhoushuguang/rose/common/net/chttp"
)

type Conf struct {
	Server *chttp.Config `yaml:"server"`
}
