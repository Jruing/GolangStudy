package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Login" method:"post" summary:"用户登录接口"`
	Username string `v:"required" dc:"用户名"`
	Password string `v:"required" dc:"密码"`
}

type LoginRes struct {
	//g.Meta `json:"g.Meta"`
	Code int64  `v:"required"`
	Msg  string `v:"required"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"Logout" method:"post" summary:"用户注销接口"`
}

type LogoutRes struct {
	g.Meta `json:"g.Meta"`
}
