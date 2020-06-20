package config

import "time"

const (
	TimeLocal     = "Asia/Shanghai"       //设置时区
	TimeTemplate1 = "2006-01-02 15:04:05" //常规类型
	TimeTemplate2 = "2006/01/02 15:04:05" //其他类型
	TimeTemplate3 = "2006-01-02"          //其他类型
	TimeTemplate4 = "15:04:05"            //其他类型
)

var (
	Local, _ = time.LoadLocation(TimeLocal)
)
