package model

import "time"

type User struct {
	Id         int64
	Name       string
	LoginName  string
	Role       string
	Pwd        string
	CreateTime time.Time
}
