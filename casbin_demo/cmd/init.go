package cmd

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

func init() {
	a, _ := gormadapter.NewAdapter("mysql", "casbin:J8R8dwYEVZbDE9VJ@tcp(mysql.sqlpub.com:3306)/casbin", true) // Your driver and data source.
	Enforcer, _ = casbin.NewEnforcer("conf/rbac_model.conf", a)
	err := Enforcer.LoadPolicy()
	if err != nil {
		return
	}
}
