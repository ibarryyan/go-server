package auth

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	a "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	enforcer *casbin.Enforcer
	adapter  *a.Adapter
)

func init() {
	var err error
	adapter, err = a.NewAdapter("mysql", "root:12345@tcp(127.0.0.1:3306)/go_app", true)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}
	m, err := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
		`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}
	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
}

func AddPolicy(role string, res string, action string) bool {
	result, err := enforcer.AddPolicy(role, res, action)
	if err != nil {
		panic(err)
	}
	return result
}

func CheckEnforce(role string, res string, action string) bool {
	result, _ := enforcer.Enforce(role, res, action)
	return result
}

func DeletePolicy(role string, res string, action string) bool {
	result, _ := enforcer.DeletePermissionForUser(role, action)
	return result
}
