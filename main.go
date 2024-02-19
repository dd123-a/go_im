package wenqianIm

import (
	"wenqianIm/conf"
	"wenqianIm/router"
	"wenqianIm/service"
)

func main() {
	conf.Init()
	go service.Manager.Start()
	r:=router.NewROuter()
	_=r.Run(conf.HttpPort)
}
