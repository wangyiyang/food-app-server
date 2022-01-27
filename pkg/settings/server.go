package settings

import "time"

type Server struct {
	AppName      string
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func ServerInit() {
	mapTo("server", ServerSetting)
}
