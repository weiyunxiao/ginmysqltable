package controller

import (
	"fmt"
	"ginmysqltable/config"
	"ginmysqltable/model"
	"github.com/gin-gonic/gin"
)

/*
应用首页
*/
func Index(c *gin.Context) {
	data := map[string]interface{}{
		"Dbname":    config.SystemConfig.Mysql.Dbname,
		"err":       "",
		"tableList": []model.Table{},
	}
	sqlStr := `show table status` //获取数据库中表记录
	var createTable []model.Table
	err := config.Db.Select(&createTable, sqlStr)
	if err != nil {
		data["err"] = fmt.Sprintf("查询数据库中的mysql表失败, err:%v\n", err)
	} else {
		data["tableList"] = &createTable
	}
	for index, v := range createTable {
		var createObj model.CreateTable
		sqlStr := `show create table ` + v.Name
		err := config.Db.Get(&createObj, sqlStr)
		if err != nil {
			data["err"] = data["err"].(string) + fmt.Sprintf("查询数据库中的mysql表失败, err:%v\n", err)
			break
		} else {
			createTable[index].TableCreateStr = createObj.CreateStr
		}
	}
	c.HTML(200, "index.html", data)
}
