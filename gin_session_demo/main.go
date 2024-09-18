package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 使用Redis存储Session
	store, err := redis.NewStore(
		16,                     // 最大空闲链接数量，过大会浪费，过小将来会触发性能瓶颈
		"tcp",                  // 指定与Redis服务器通信的网络类型，通常为"tcp"
		"192.168.198.129:6379", // Redis服务器的地址，格式为"host:port"
		"123456",               // Redis服务器的密码，如果没有密码可以为空字符串
		[]byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm"), // authentication key
		[]byte("o6idlo2cb9f9pb6h46fimllw481ldebi"), // encryption key
	)
	// 使用内存
	//store := memstore.NewStore([]byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm"), []byte("o6idlo2cb9f9pb6h46fimllw481ldebi"))
	router.Use(sessions.Sessions("sessionId", store))

	router.GET("/login", func(context *gin.Context) {
		session := sessions.Default(context)
		session.ID()
		session.Set("username", "Jruing")
		err := session.Save()
		if err != nil {
			return
		}
		context.JSON(200, gin.H{
			"Code": 1,
		})
	})
	r := router.Group("/api")
	//r.Use(middleware.IsLogin)

	r.GET("/get", func(context *gin.Context) {
		session := sessions.Default(context)
		context.JSON(200, gin.H{
			"Code":     1,
			"Msg":      session.ID(),
			"Username": session.Get("username"),
		})
	})
	err = router.Run(":8000")
	if err != nil {
		return
	}
}
