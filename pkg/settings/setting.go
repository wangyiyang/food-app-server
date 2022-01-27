package settings

import (
	"log"

	"gopkg.in/ini.v1"
)

var cfg *ini.File

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func Setup() {
	var err error
	cfg, err = ini.Load("./config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse './config/app.ini': %v", err)
	}

}
