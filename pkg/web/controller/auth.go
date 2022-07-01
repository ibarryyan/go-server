package controller

import (
	"count_num/pkg/web/auth"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type AuthResp struct {
	Role     string `json:"role"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func (auth AuthResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"role":     auth.Role,
		"resource": auth.Resource,
		"action":   auth.Action,
	})
}

func AddPolicy(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	a := AuthResp{}
	json.Unmarshal(bytes, &a)
	if err != nil {
		panic(err)
	}
	res := auth.AddPolicy(a.Role, a.Resource, a.Action)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}
