# Gin web框架的范例代码


## 访问路径
http://localhost:8080/ping

## 启动代码
入口源码: quick-start.go
启动: go run quick-start.go


## 编译
* 一般编译
go build quick-start.go

* 交叉编译-Windows下输出到其它系统
```bat
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build quick-start.go
```


* 交叉编译-Linux下输出其它系统
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build quick-start.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build quick-start.go

* 交叉编译依赖 golang.org/x下包的问题解决
https://shockerli.net/post/go-get-golang-org-x-solution/

设置 GOPROXY 环境变量
Unix下：	export GOPROXY=https://goproxy.io

Windows cmd:			SET GOPROXY=https://goproxy.io
Windows powershell：	$env:GOPROXY = "https://goproxy.io"
