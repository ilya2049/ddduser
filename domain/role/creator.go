package role

import (
	"context"
	"errors"
)

var ErrRoleWithSameLevelAlreadyExists = errors.New("role with the same level already exists")

func NewCreator(roleRepository Repository) *Creator {
	return &Creator{
		roleRepository: roleRepository,
	}
}

type Creator struct {
	roleRepository Repository
}

func (f *Creator) CreateTestRoles(ctx context.Context) error {
	roles := []Role{
		newTestAdmin(),
		newTestModerator(),
		newTestGuest(),
	}

	for _, role := range roles {
		if err := f.CreateRole(ctx, role); err != nil {
			return err
		}
	}

	return nil
}

func (f *Creator) CreateRole(ctx context.Context, rl Role) error {
	ok, err := f.roleRepository.HasRoleWithLevel(ctx, rl.Level())
	if err != nil {
		return err
	}

	if ok {
		return ErrRoleWithSameLevelAlreadyExists
	}

	_, err = f.roleRepository.Add(ctx, rl)

	if err != nil {
		return err
	}

	return nil
}
