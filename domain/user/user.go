package user

import (
	"context"
	"ddduser/domain/auth"
	"errors"
)

var (
	ErrUserDoesNotExist = errors.New("user does not exist")
)

func NewAdmin(
	credentials Credentials,
	roleRepository RoleRepository,
) (User, error) {
	return newUser(credentials, auth.RoleLevelAdmin, roleRepository)
}

func newUser(
	credentials Credentials,
	roleLevel auth.RoleLevel,
	roleRepository RoleRepository,
) (User, error) {
	role, err := roleRepository.GetByLevel(context.TODO(), roleLevel)
	if err != nil {
		return User{}, err
	}

	return User{
		credentials: credentials,
		role:        role,
	}, nil
}

type User struct {
	id          ID
	ownerID     ID
	credentials Credentials
	role        Role
}

func (u User) NewChildUser(
	credentials Credentials,
	roleLevel auth.RoleLevel,
	roleRepository RoleRepository,
) (User, error) {
	u, err := newUser(credentials, roleLevel, roleRepository)

	if err != nil {
		return User{}, err
	}

	u.ownerID = u.id

	return u, nil
}

func (u User) ID() ID {
	return u.id
}

func (u User) Credentials() Credentials {
	return u.credentials
}

func (u User) Role() Role {
	return u.role
}

func (u *User) Identify(id ID) {
	u.id = id
}

func (u *User) UpdateCredentials(newCredentials Credentials) {
	u.credentials = newCredentials
}

func (u User) IsOwnedBy(other User) bool {
	return u.ownerID == other.id
}

func (u User) Is(other User) bool {
	return u.id == other.id
}

type ID = int

type Repository interface {
	Add(context.Context, User) (ID, error)
	Update(context.Context, User) error
	Delete(context.Context, ID) error
	Get(context.Context, ID) (User, error)
	GetAdmin(context.Context) (User, error)
	List(context.Context) ([]User, error)
}
