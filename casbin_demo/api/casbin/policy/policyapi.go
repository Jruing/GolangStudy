package policy

import (
	"casbin_demo/cmd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Sub string
	Dom string
	Obj string
	Act string
}

func (receiver Request) PolicyAdd(c *gin.Context) {
	enforcer := cmd.Enforcer
	var request Request
	err := c.Bind(&request)
	if err != nil {
		return
	}
	sub := request.Sub
	dom := request.Dom
	obj := request.Obj
	act := request.Act
	policy := []string{sub, dom, obj, act}
	_, err = enforcer.AddPolicy(policy)
	if err != nil {
		return
	}
}

func (receiver Request) PolicyDel(c *gin.Context) {
	enforcer := cmd.Enforcer
	var request Request
	err := c.Bind(&request)
	if err != nil {
		return
	}
	sub := request.Sub
	dom := request.Dom
	obj := request.Obj
	act := request.Act
	fmt.Println(sub, dom, obj, act)
	_, err = enforcer.RemovePolicy(sub, dom, obj, act)
	if err != nil {
		fmt.Println(err)
	}
}

func (receiver Request) PolicyMatch(c *gin.Context) {
	enforcer := cmd.Enforcer
	var request Request
	err := c.Bind(&request)
	if err != nil {
		return
	}
	sub := request.Sub
	dom := request.Dom
	obj := request.Obj
	act := request.Act
	// 校验权限.
	_, err = enforcer.Enforce(sub, dom, obj, act)
	if err != nil {
		fmt.Println("权限校验失败")
		return
	}
}
