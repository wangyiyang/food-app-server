package settings

// Log 日志配置
type Log struct {
	BasePath string
	SubDir   bool
}

// LogSetting 日志配置结构体
var LogSetting = &Log{}

func LogInit() {
	mapTo("log", LogSetting)

}
