package main

import (
	"ginmysqltable/config"
	"ginmysqltable/controller"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func main() {
	//设置gin 运行模式
	gin.SetMode(config.SystemConfig.GinMode)
	r := gin.Default()
	//创建模板中需要使用的函数
	r.SetFuncMap(template.FuncMap{
		"timeStr": timeStr,
	})
	r.Static("/static", "./static")
	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.Index)
	r.Run(":9000")
}

func timeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
