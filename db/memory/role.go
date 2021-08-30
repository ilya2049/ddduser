package memory

import (
	"context"
	"ddduser/domain/auth"
	"ddduser/domain/role"
	"ddduser/domain/user"
)

type RoleRepository struct {
	lastID int

	roles []role.Role
}

func (r *RoleRepository) Add(_ context.Context, rl role.Role) (role.ID, error) {
	r.lastID++

	rl.Identify(r.lastID)

	r.roles = append(r.roles, rl)

	return rl.ID(), nil
}

func (r *RoleRepository) GetByLevel(_ context.Context, roleLevel auth.RoleLevel) (user.Role, error) {
	for _, rl := range r.roles {
		if rl.Level() == roleLevel {
			return user.NewRole(rl.ID(), rl.Level()), nil
		}
	}

	return user.Role{}, role.ErrRoleDoesNotExist
}

func (r *RoleRepository) HasRoleWithLevel(_ context.Context, roleLevel auth.RoleLevel) (bool, error) {
	for _, rl := range r.roles {
		if rl.Level() == roleLevel {
			return true, nil
		}
	}

	return false, nil
}
