package role

import (
	"ddduser/domain/auth"
	"errors"
)

var (
	ErrRoleDoesNotExist = errors.New("role does not exist")
)

func New(name string, level auth.RoleLevel) Role {
	return Role{
		name:  name,
		level: level,
	}
}

type Role struct {
	id    ID
	name  string
	level auth.RoleLevel
}

type ID = int

func (r *Role) Identify(id ID) {
	r.id = id
}

func (r Role) Level() auth.RoleLevel {
	return r.level
}

func (r Role) Name() string {
	return r.name
}

func (r Role) ID() ID {
	return r.id
}

type Repository interface {
	Add(Role) (ID, error)
}
