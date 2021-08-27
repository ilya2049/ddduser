package user

import "ddduser/domain/auth"

type ModeratorContext struct {
	moderator User

	roleRepository RoleRepository
}

func (mc *ModeratorContext) NewUser(
	name string,
	roleLevel auth.RoleLevel,
) (User, error) {
	if roleLevel == auth.RoleLevelGuest {
		return mc.moderator.NewChildUser(name, roleLevel, mc.roleRepository)
	}

	return User{}, ErrOperationIsForbiddenForCurrentUser
}

func (mc *ModeratorContext) CanUpdateUser(u User) error {
	if u.Is(mc.moderator) || u.IsOwnedBy(mc.moderator) {
		return nil
	}

	return ErrOperationIsForbiddenForCurrentUser
}

func (mc *ModeratorContext) CanDeleteUser(u User) error {
	if u.IsOwnedBy(mc.moderator) {
		return nil
	}

	return ErrOperationIsForbiddenForCurrentUser
}
