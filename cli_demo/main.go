package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var config string
var output string

// 初始化根命令
var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "这是一个简单的描述",
	Long:  "这是一个详细的描述",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("这是一个Cobra示例", config)
	},
}

// 创建一个命令echo
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "输出当前服务信息",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("当前服务地址：http://127.0.0.0:8000")
	}}

// 创建有一个带位置参数的命令init
var initCmd = &cobra.Command{
	Use:   "init <name>",
	Short: "初始化一个指定服务的配置",
	Args:  cobra.ExactArgs(1), // 解析位置参数name
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0] // 获取位置参数name的值
		fmt.Println("开始执行初始化", name)
	},
}

// 创建命令server
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "获取服务列表",
}

// 创建命令ls，用于获取服务列表明细
var serverListCmd = &cobra.Command{
	Use:   "ls",
	Short: "获取服务列表",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("当前服务列表：\naa\nbb\ncc")
		fmt.Println("将文件输出到", output)
	},
}

// 创建带有位置参数的命令get，用于获取具体服务的信息
var serverGetCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "获取具体服务信息",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		servername := args[0]
		fmt.Println("获取服务信息：", servername)
	},
}

// 运行初始化配置函数
func initConfig() {
	fmt.Println("运行初始化")
}

func main() {
	// 加载初始化配置
	cobra.OnInitialize(initConfig)

	// 将echo，init，server这三个命令加入到根命令下
	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(serverCmd)

	// 将serverListCmd serverGetCmd 这两个命令作为子命令加入到server命令下
	serverCmd.AddCommand(serverListCmd)
	serverCmd.AddCommand(serverGetCmd)

	// Persistent Flags，持久化的flag，全局可以使用，可以给任一命令添加选项
	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.firstappname.yaml)")

	// Flags 局部使用，为某一个命令添加选项
	serverListCmd.Flags().StringVar(&output, "output", "", "输出文件")
	serverGetCmd.Flags().StringVar(&output, "output", "", "输出文件")
	// 把 output 设置为必选项
	err := serverListCmd.MarkFlagRequired("output")
	if err != nil {
		return
	}
	// 运行
	if err := rootCmd.Execute(); err != nil {
		return
	}
}
