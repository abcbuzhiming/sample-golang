# Gin web��ܵķ�������


## ����·��
http://localhost:8080/ping

## ��������
���Դ��: quick-start.go
����: go run quick-start.go


## ����
* һ�����
go build quick-start.go

* �������-Windows�����������ϵͳ
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


* �������-Linux���������ϵͳ
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build quick-start.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build quick-start.go

* ����������� golang.org/x�°���������
https://shockerli.net/post/go-get-golang-org-x-solution/

���� GOPROXY ��������
Unix�£�	export GOPROXY=https://goproxy.io

Windows cmd:			SET GOPROXY=https://goproxy.io
Windows powershell��	$env:GOPROXY = "https://goproxy.io"
