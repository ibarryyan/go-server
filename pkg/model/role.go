package model

import "encoding/json"

type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (role Role) TableName() string {
	return "role_info"
}

func (role Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   role.Id,
		"name": role.Name,
	})
}
