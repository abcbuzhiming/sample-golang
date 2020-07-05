# goframe框架的范例代码

## 目录说明
api	业务接口	接收/解析用户输入参数的入口/接口层。
model	数据模型	数据管理层，仅用于操作管理数据，如数据库操作。
service	逻辑封装	业务逻辑封装层，实现特定的业务需求，可供不同的包调用。
boot	初始化包	用于项目初始化参数设置，往往作为main.go中第一个被import的包。
config	配置管理	所有的配置文件存放目录。
docker	镜像文件	Docker镜像相关依赖文件，脚本文件等等。
document	项目文档	Documentation项目文档，如: 设计文档、帮助文档等等。
i18n	I18N国际化	I18N国际化配置文件目录。
library	公共库包	公共的功能封装包，往往不包含业务需求实现。
packed	打包目录	将资源文件打包的Go文件存放在这里，boot包初始化时会自动调用。
public	静态目录	仅有该目录下的文件才能对外提供静态服务访问。
router	路由注册	用于路由统一的注册管理。
template	模板文件	MVC模板文件存放的目录。
vendor	第三方包	第三方依赖包存放目录（可选, 未来会被淘汰）。
Dockerfile	镜像描述	云原生时代用于编译生成Docker镜像的描述文件。
go.mod	依赖管理	使用Go Module包管理的依赖描述文件。
main.go	入口文件	程序入口文件。


## 设置 GOPROXY 环境变量
Unix下：	export GOPROXY=https://goproxy.io

Windows cmd:			SET GOPROXY=https://goproxy.io
Windows powershell：	$env:GOPROXY = "https://goproxy.io"

## 启动代码
入口源码: main.go
启动: go run main.go


