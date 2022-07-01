package controller

import "count_num/pkg/dao/impl"

type UserControllerImpl struct {
	dao *impl.UserDaoImpl
}

type UserController interface {
}
