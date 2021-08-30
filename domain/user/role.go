package user

import (
	"context"
	"ddduser/domain/auth"
)

func NewRole(id int, level auth.RoleLevel) Role {
	return Role{
		id:    id,
		level: level,
	}
}

type Role struct {
	id    int
	level auth.RoleLevel
}

func (r Role) Level() auth.RoleLevel {
	return r.level
}

type RoleRepository interface {
	GetByLevel(context.Context, auth.RoleLevel) (Role, error)
}
