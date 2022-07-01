package model

import "time"

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	LoginName  string    `json:"login_name"`
	Role       int64     `json:"role"`
	Pwd        string    `json:"pwd"`
	CreateTime time.Time `json:"create_time"`
}
