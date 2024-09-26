package main

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os/exec"
)

//go:embed static/*
var data string

func main() {
	r := gin.Default()
	r.Static("../static", "static")
	r.LoadHTMLGlob("static/*.tmpl")
	r.Use()
	r.GET("/install", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "install.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.POST("/inits", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Code": 2,
			"Msg":  "msg",
		})
	})
	// 重启项目
	r.GET("/restart", func(ctx *gin.Context) {
		cmd := exec.Command("killall", "-HUP", "appweiyigeek")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error executing restart command:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restart Gin server."})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Gin server restarted successfully."})
	})
	user := r.Group("/user")
	{
		user.POST("")
	}
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
