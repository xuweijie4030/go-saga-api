package daemon

import (
	"gosagaapi/app/global"
	"gosagaapi/app/logic"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/xhttp"
)

type InitGlobalDaemon struct {
}

func NewInitGlobalDaemon() *InitGlobalDaemon {
	return &InitGlobalDaemon{}
}

func (d *InitGlobalDaemon) Handle() {
	d.initJwt()
}

func (d *InitGlobalDaemon) initJwt() {
	global.Jwt = xhttp.NewXJwt(xhttp.JwtOptions{
		Secret: config.GetString("Jwt.Secret"),
		Lookup: config.GetString("Jwt.Lookup"),
		Auth:   logic.NewAuthLogic(),
	})
}
