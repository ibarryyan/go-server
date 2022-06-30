package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"log"
)

func Main() {
	// 使用 MySQL 数据库初始化一个 Xorm 适配器
	a, err := xormadapter.NewAdapter("mysql", "root:12345@tcp(127.0.0.1:3306)/go_app", true)
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

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	e.AddPolicy("sub", "resource", "write")

	res, err := e.Enforce("sub", "resource", "write")

	fmt.Println(res)

}
