package user

import (
	"ddduser/domain/auth"

	"errors"
)

var ErrOperationIsForbiddenForCurrentUser = errors.New("the operation is forbidden for the current user")

type CurrentUser interface {
	NewUser(Credentials, auth.RoleLevel) (User, error)
	CanUpdateUser(User) error
	CanDeleteUser(User) error
	CanReadUser(User) error
}

func NewCurrentUserFactory(roleRepository RoleRepository) *CurrentUserFactory {
	return &CurrentUserFactory{
		roleRepository: roleRepository,
	}
}

type CurrentUserFactory struct {
	roleRepository RoleRepository
}

func (f *CurrentUserFactory) NewCurrentUser(u User) CurrentUser {
	switch u.Role().Level() {
	case auth.RoleLevelAdmin:
		return &AdminContext{
			admin: u,

			roleRepository: f.roleRepository,
		}

	case auth.RoleLevelModerator:
		return &ModeratorContext{
			moderator: u,

			roleRepository: f.roleRepository,
		}

	case auth.RoleLevelGuest:
		return &GuestContext{
			guest: u,
		}

	default:
		panic("role of the current user is unknown")
	}
}
