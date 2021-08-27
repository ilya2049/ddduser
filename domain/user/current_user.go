package user

import (
	"ddduser/domain/auth"

	"errors"
)

var ErrOperationIsForbiddenForCurrentUser = errors.New("the operation is forbidden for the current user")

type CurrentUser interface {
	NewUser(string, auth.RoleLevel) (User, error)
	CanUpdateUser(User) error
	CanDeleteUser(User) error
}

func NewCurrentUserFactory(roleRepository RoleRepository) *CurrentUserFactory {
	return &CurrentUserFactory{
		roleRepository: roleRepository,
	}
}

type CurrentUserFactory struct {
	roleRepository RoleRepository
}

var ErrUnknownRoleOfCurrentUser = errors.New("role of the current user is unknown")

func (f *CurrentUserFactory) NewCurrentUser(u User) (CurrentUser, error) {
	switch u.Role().Level() {
	case auth.RoleLevelAdmin:
		return &AdminContext{
			admin: u,

			roleRepository: f.roleRepository,
		}, nil

	case auth.RoleLevelModerator:
		return &ModeratorContext{
			moderator: u,

			roleRepository: f.roleRepository,
		}, nil

	case auth.RoleLevelGuest:
		return &GuestContext{
			guest: u,
		}, nil

	default:
		return nil, ErrUnknownRoleOfCurrentUser
	}
}
