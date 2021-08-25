package user

import (
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

func (ac *AdminCreator) NewAdmin(name string) (User, error) {
	admin, err := New(name, auth.RoleLevelAdmin, ac.roleRepository)
	if err != nil {
		return User{}, err
	}

	_, err = ac.userRepository.GetAdmin()

	if err == nil {
		return User{}, ErrOnlyOneAdmin
	}

	if errors.Is(err, ErrUserDoesNotExist) {
		adminID, err := ac.userRepository.Add(admin)
		if err != nil {
			return User{}, err
		}

		admin.Identify(adminID)

		return admin, nil
	}

	return User{}, err
}
