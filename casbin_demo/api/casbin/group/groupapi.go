package group

import (
	"casbin_demo/cmd"
	"github.com/gin-gonic/gin"
)

type Request struct {
	User string
	Role string
	Dom  string
}

func GroupAdd(c *gin.Context) {
	enforcer := cmd.Enforcer
	var request Request
	err := c.Bind(&request)
	if err != nil {
		return
	}
	user := request.User
	dom := request.Dom
	role := request.Role
	_, err = enforcer.AddRoleForUserInDomain(user, role, dom)
	if err != nil {
		return
	}
}

func GroupDel(c *gin.Context) {
	enforcer := cmd.Enforcer
	var request Request
	err := c.Bind(&request)
	if err != nil {
		return
	}
	user := request.User
	dom := request.Dom
	role := request.Role
	_, err = enforcer.DeleteRoleForUserInDomain(user, role, dom)
	if err != nil {
		return
	}
}
