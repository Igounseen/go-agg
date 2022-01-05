package main

import (
	"hello/lg"
)

func main() {
	// console
	consoleProfile := lg.NewProfile("console", lg.Info)
	consoleAppender := lg.NewConsoleAppender()
	consoleProfile.AddAppender(consoleAppender)

	// file

	mylg := lg.NewMylg()
	mylg.AddProfile(consoleProfile)

	mylg.Debug("测试日志鸭debug")
	mylg.Info("测试日志鸭info")
	mylg.Error("测试日志鸭error")

}
