package user

import (
    //"github.com/gogf/gf-demos/app/service/user"
    "github.com/gogf/gf-demos/library/response"
    "github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
type Controller struct{}


// 登录请求参数，用于前后端交互参数格式约定
type SignInRequest struct {
    UserName string `v:"required#账号不能为空"`
    Password string `v:"required#密码不能为空"`
}

// 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
    var data *SignInRequest
    if err := r.Parse(&data); err != nil {
        response.JsonExit(r, 1, err.Error())
	}
	/*
    if err := user.SignIn(data.Passport, data.Password, r.Session); err != nil {
        response.JsonExit(r, 1, err.Error())
    } else {
        response.JsonExit(r, 0, "ok")
	}
	*/
	response.JsonExit(r, 0, "ok")
}