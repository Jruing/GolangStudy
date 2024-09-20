package login

import (
	"context"
	"fmt"
	"gf_demo/api/login/v1"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	fmt.Println("=============", req.Username)
	a, _ := g.RequestFromCtx(ctx).GetJson()
	fmt.Println(a.Get("Username"))
	g.RequestFromCtx(ctx).Response.Header().Set("Content-Type", "application/json")
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"id":   1,
		"name": "john",
	})
	return
	//return nil, gerror.NewCode(gcode.CodeOK)
}
