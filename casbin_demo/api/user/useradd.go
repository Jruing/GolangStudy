package user

import (
	"casbin_demo/model"
	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return
	}

}
