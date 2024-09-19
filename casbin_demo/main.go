package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a, _ := gormadapter.NewAdapter("mysql", "casbin:J8R8dwYEVZbDE9VJ@tcp(mysql.sqlpub.com:3306)/casbin", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("conf/rbac_model.conf", a)

	// 加载策略
	err := e.LoadPolicy()
	if err != nil {
		return
	}

	//// 校验权限.
	//_, err = e.Enforce("alice", "data1", "read")
	//if err != nil {
	//	fmt.Println("权限校验失败")
	//	return
	//}
	//fmt.Println("权限校验")
	var policy [][]string
	slice1 := []string{"运维", "百度", "CMDB", "read"}

	policy = append(policy, slice1)
	e.AddPolicies(policy)
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	//e.AddRoleForUser("张三", "运维", "百度")
	//e.AddRoleForUser("李四", "运维", "百度")
	//e.AddRoleForUser("王五", "开发", "百度")
	// 修改策略
	//var c []string
	//c = append(c, "aaa")
	//e.UpdatePolicy(c, c)

	// 保存策略
	err = e.SavePolicy()
	if err != nil {
		return
	}
	fmt.Println("完成")
}
