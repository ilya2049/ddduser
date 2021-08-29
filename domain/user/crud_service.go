package user

import "context"

type CRUDService struct {
	userRepository Repository
}

func (s *CRUDService) CreateUser(ctx context.Context, u User) (ID, error) {
	return s.userRepository.Add(u)
}

func (s *CRUDService) UpdateUser(ctx context.Context, cu CurrentUser, u User) error {
	if err := cu.CanUpdateUser(u); err != nil {
		return err
	}

	return s.userRepository.Update(u)
}

func (s *CRUDService) DeleteUser(ctx context.Context, cu CurrentUser, id ID) error {
	u, err := s.userRepository.Get(id)
	if err != nil {
		return err
	}

	if err := cu.CanDeleteUser(u); err != nil {
		return err
	}

	return s.userRepository.Delete(u.ID())
}

func (s *CRUDService) GetUser(ctx context.Context, cu CurrentUser, id ID) (User, error) {
	u, err := s.userRepository.Get(id)
	if err != nil {
		return User{}, err
	}

	if err := cu.CanReadUser(u); err != nil {
		return User{}, err
	}

	return u, nil
}
