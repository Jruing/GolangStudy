package cmd

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os/exec"
)

var Enforcer *casbin.Enforcer

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "Casbin 权限管理系统",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Casbin 权限管理系统")
	},
}

var serverCmd = &cobra.Command{
	Use:   "Server",
	Short: "启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.Static("/static", "../static")
		r.Use()
		r.GET("/install", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Main website",
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
	},
}

func init() {

	// 初始化casbin适配器
	adapter, _ := gormadapter.NewAdapter("mysql", "casbin:J8R8dwYEVZbDE9VJ@tcp(mysql.sqlpub.com:3306)/casbin", true) // Your driver and data source.
	Enforcer, _ = casbin.NewEnforcer("conf/rbac_model.conf", adapter)
	err := Enforcer.LoadPolicy()
	if err != nil {
		return
	}

	rootCmd.AddCommand(serverCmd)

}
