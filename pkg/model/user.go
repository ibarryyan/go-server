package model

import (
	"encoding/json"
	"time"
)

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	LoginName  string    `json:"login_name"`
	Role       int64     `json:"role"`
	Pwd        string    `json:"pwd"`
	CreateTime time.Time `json:"create_time"`
}

func (user User) TableName() string {
	return "user_info"
}

func (user User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":          user.Id,
		"name":        user.Name,
		"login_name":  user.LoginName,
		"role":        user.Role,
		"pwd":         user.Pwd,
		"create_time": user.CreateTime,
	})
}

//Redis类似序列化操作
func (user User) MarshalBinary() ([]byte, error) {
	return json.Marshal(user)
}

func (user User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &user)
}
