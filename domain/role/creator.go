package role

import "errors"

var ErrRoleWithSameLevelAlreadyExists = errors.New("role with the same level already exists")

func NewCreator(roleRepository Repository) *Creator {
	return &Creator{
		roleRepository: roleRepository,
	}
}

type Creator struct {
	roleRepository Repository
}

func (f *Creator) CreateRole(rl Role) error {
	ok, err := f.roleRepository.HasRoleWithLevel(rl.Level())
	if err != nil {
		return err
	}

	if ok {
		return ErrRoleWithSameLevelAlreadyExists
	}

	_, err = f.roleRepository.Add(rl)

	if err != nil {
		return err
	}

	return nil
}