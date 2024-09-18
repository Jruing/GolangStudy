package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a, _ := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(192.168.198.129:3306)/casbin_authority", true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("conf/rbac_model.conf", a)

	// 加载策略
	err := e.LoadPolicy()
	if err != nil {
		return
	}

	// 校验权限.
	_, err = e.Enforce("alice", "data1", "read")
	if err != nil {
		return
	}
	var policy [][]string
	slice1 := []string{"1", "1", "1", "1"}

	policy = append(policy, slice1)
	e.AddPolicies(policy)
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// 修改策略
	var c []string
	c = append(c, "aaa")
	e.UpdatePolicy(c, c)

	// 保存策略
	err = e.SavePolicy()
	if err != nil {
		return
	}
}
