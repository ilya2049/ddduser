package memory

import (
	"context"
	"ddduser/domain/auth"
	"ddduser/domain/user"
)

type UserRepository struct {
	lastID int

	users []user.User
}

func (r *UserRepository) Add(_ context.Context, u user.User) (user.ID, error) {
	r.lastID++

	u.Identify(r.lastID)

	r.users = append(r.users, u)

	return u.ID(), nil
}

func (r *UserRepository) Update(_ context.Context, u user.User) error {
	var userIdx = -1

	for i, usr := range r.users {
		if usr.ID() == u.ID() {
			userIdx = i

			break
		}
	}

	if userIdx == -1 {
		return user.ErrUserDoesNotExist
	}

	r.users[userIdx] = u

	return nil
}

func (r *UserRepository) Delete(_ context.Context, id user.ID) error {
	var userIdx = -1

	for i, usr := range r.users {
		if usr.ID() == id {
			userIdx = i

			break
		}
	}

	if userIdx == -1 {
		return user.ErrUserDoesNotExist
	}

	lastUserIdx := len(r.users) - 1

	r.users[userIdx] = r.users[lastUserIdx]
	r.users[lastUserIdx] = user.User{}
	r.users = r.users[:lastUserIdx]

	return nil
}

func (r *UserRepository) Get(_ context.Context, id user.ID) (user.User, error) {
	for _, usr := range r.users {
		if usr.ID() == id {
			return usr, nil
		}
	}

	return user.User{}, user.ErrUserDoesNotExist
}

func (r *UserRepository) GetAdmin(_ context.Context) (user.User, error) {
	for _, usr := range r.users {
		if usr.Role().Level() == auth.RoleLevelAdmin {
			return usr, nil
		}
	}

	return user.User{}, user.ErrUserDoesNotExist
}

func (r *UserRepository) List(_ context.Context) ([]user.User, error) {
	return append(make([]user.User, 0, len(r.users)), r.users...), nil
}
