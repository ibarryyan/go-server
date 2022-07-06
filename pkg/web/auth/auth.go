package auth

import (
	"context"
	"count_num/pkg/config"
	model2 "count_num/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	a "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	enforcer     *casbin.Enforcer
	adapter      *a.Adapter
	TokenTimeOut = time.Second * 3600
	RDB          = config.RDB
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
		m = my_func(r.sub,p.sub,r.obj,p.obj)
		`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}
	//自定义匹配函数
	enforcer, err = casbin.NewEnforcer(m, adapter)
	enforcer.AddFunction("my_func", KeyMatchFunc)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
}

//自定义匹配函数
//规则： (r.sub == p.sub && r.obj == p.obj) || (r.sub == p.sub && p.obj == '*')
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	KeyMatch := func(rsub, psub, robj, pobj string) bool {
		return (rsub == psub && robj == pobj) || (rsub == psub && pobj == "*")
	}
	return (bool)(KeyMatch(args[0].(string), args[1].(string), args[2].(string), args[3].(string))), nil
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

func SetToken(ctx context.Context, token string, user model2.User) bool {
	RDB.Set(ctx, token, user, TokenTimeOut)
	return true
}

func GetToken(ctx context.Context, token string) model2.User {
	var user model2.User
	u, err := RDB.Get(ctx, token).Result()
	if len(u) == 0 {
		return user
	}
	err = json.Unmarshal([]byte(u), &user)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	return user
}
