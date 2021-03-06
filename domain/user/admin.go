package user

import (
	"context"
	"ddduser/domain/auth"
	"errors"
)

var ErrOnlyOneAdmin = errors.New("there can be only one admin")

func NewAdminCreator(
	roleRepository RoleRepository,
	userRepository Repository,
) *AdminCreator {
	return &AdminCreator{
		roleRepository: roleRepository,
		userRepository: userRepository,
	}
}

type AdminCreator struct {
	roleRepository RoleRepository
	userRepository Repository
}

func (ac *AdminCreator) CreateAdmin(ctx context.Context, credentials Credentials) (User, error) {
	admin, err := NewAdmin(credentials, ac.roleRepository)
	if err != nil {
		return User{}, err
	}

	_, err = ac.userRepository.GetAdmin(ctx)

	if err == nil {
		return User{}, ErrOnlyOneAdmin
	}

	if errors.Is(err, ErrUserDoesNotExist) {
		adminID, err := ac.userRepository.Add(ctx, admin)
		if err != nil {
			return User{}, err
		}

		admin.Identify(adminID)

		return admin, nil
	}

	return User{}, err
}

type AdminContext struct {
	admin User

	roleRepository RoleRepository
}

func (ac *AdminContext) NewUser(
	credentials Credentials,
	roleLevel auth.RoleLevel,
) (User, error) {
	if roleLevel == auth.RoleLevelModerator {
		return ac.admin.NewChildUser(credentials, roleLevel, ac.roleRepository)
	}

	return User{}, ErrOperationIsForbiddenForCurrentUser
}

func (ac *AdminContext) CanUpdateUser(u User) error {
	if u.Role().Level() == auth.RoleLevelModerator || u.Is(ac.admin) {
		return nil
	}

	return ErrOperationIsForbiddenForCurrentUser
}

func (ac *AdminContext) CanDeleteUser(u User) error {
	if u.Role().Level() == auth.RoleLevelModerator {
		return nil
	}

	return ErrOperationIsForbiddenForCurrentUser
}

func (ac *AdminContext) CanReadUser(_ User) error {
	return nil
}
