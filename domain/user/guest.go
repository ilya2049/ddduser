package user

import "ddduser/domain/auth"

type GuestContext struct {
	guest User
}

func (gc *GuestContext) NewUser(
	name string,
	roleLevel auth.RoleLevel,
) (User, error) {
	return User{}, ErrOperationIsForbiddenForCurrentUser
}

func (gc *GuestContext) CanUpdateUser(u User) error {
	if u.Is(gc.guest) {
		return nil
	}

	return ErrOperationIsForbiddenForCurrentUser
}

func (gc *GuestContext) CanDeleteUser(u User) error {
	if u.Is(gc.guest) {
		return nil
	}

	return ErrOperationIsForbiddenForCurrentUser
}

func (gc *GuestContext) CanReadUser(u User) error {
	return gc.CanDeleteUser(u)
}
