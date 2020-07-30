package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

//配置文件路径
const configFile = "./config/config.yaml"

//定义公开数据库连接池变量
var Db *sqlx.DB

//定义配置对就看结构体变量Auto_increment
var SystemConfig System

type System struct {
	GinMode string
	GinPort string
	Mysql   Mysql
}

type Mysql struct {
	Username     string //mysql 用户名
	Password     string //mysql 密码
	IpPort       string //mysql ip与端口
	Dbname       string //数据库名
	Config       string //其它额外的配置信息
	MaxIdleConns int    //连接最大空闲数
	MaxOpenConns int    //连接最大打开数
}

//最先加载
func init() {
	initConfig()
	initDB()
}

//获取配置相关参数
func initConfig() {
	v := viper.New()
	v.SetConfigFile(configFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败，请检查配置文件: %s \n", err))
	}
	if err := v.Unmarshal(&SystemConfig); err != nil {
		panic(fmt.Errorf("配置文件绑定到结构体失败，请检查: %s \n", err))
	}
}

//初始化数据库
func initDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		SystemConfig.Mysql.Username,
		SystemConfig.Mysql.Password,
		SystemConfig.Mysql.IpPort,
		SystemConfig.Mysql.Dbname,
		SystemConfig.Mysql.Config,
	)
	// 也可以使用MustConnect连接不成功就panic
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("数据库连接失败:%s", err))
	}
	Db.SetMaxOpenConns(SystemConfig.Mysql.MaxOpenConns)
	Db.SetMaxIdleConns(SystemConfig.Mysql.MaxIdleConns)
	return
}
