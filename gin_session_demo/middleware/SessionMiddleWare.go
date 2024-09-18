package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsLogin(context *gin.Context) {
	session := sessions.Default(context)
	if session.ID() == "" {
		//context.JSON(200, gin.H{
		//	"Code": 0,
		//	"Msg":  "无权限",
		//})
		context.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8000/login")
	}
}
