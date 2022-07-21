package logger

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var Recorder *logs.BeeLogger

func init() {
	// 设置日志生成器
	log := logs.NewLogger()
	// 读取配置
	logPath, err := beego.AppConfig.String("logpath")
	if err != nil {
		fmt.Println("load log path config error.")
	}
	logFile, err := beego.AppConfig.String("logfile")
	if err != nil {
		fmt.Println("load log filename config error.")
	}
	// todo: more options in .conf
	logConfig := fmt.Sprintf(
		`{"filename":"%s","level":7,"daily":true,"maxdays":10,"color":true}`,
		logPath+logFile,
	)
	log.SetLogger(
		logs.AdapterFile,
		logConfig,
	)
	log.SetLogger(logs.AdapterConsole)
	log.EnableFuncCallDepth(true)
	log.SetLevel(logs.LevelDebug)

	Recorder = log
}
