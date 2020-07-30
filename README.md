Go语言实现mysql表信息查看的小应用
===============
![](http://s.gincms.com/img/2020/0728/mysqltable-readme.jpg) 

#### 介绍
开发工作中，根据业务逻辑的开发，经常需要查看数据库表结构与索引等等相关信息，此应用方便程序员查看mysql表的相关信息
> 后期版本考虑实现Excel导出表信息功能

#### 相关技术
1.  Go语言
2.  Go语言框架:gin
3.  界面基于layui:https://www.layui.com/

#### 目录结构

初始的目录结构如下：

~~~
├─config                应用配置目录
├─controller            控制器(逻辑代码暂时全写到此目录下)
├─model                 与表对应的模型文件
├─static                静态文件(css/js)
├─view                  视图文件
~~~

#### 安装教程
```cassandraql
git clone https://github.com/weiyunxiao/ginmysqltable
cd ginmysqltable
#修改config目录下的config.yaml配置文件
#进入main.go文件同级目录
go run main.go
#打开浏览器输入应用url，例如:http://127.0.0.1:9000
```

更多信息 [http://www.gincms.com/](http://www.gitcms.com)