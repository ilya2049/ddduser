package user

import (
	"ddduser/domain/auth"
	"errors"
)

var (
	ErrUserDoesNotExist = errors.New("user does not exist")
)

func NewAdmin(
	name string,
	roleRepository RoleRepository,
) (User, error) {
	return newUser(name, auth.RoleLevelAdmin, roleRepository)
}

func newUser(
	name string,
	roleLevel auth.RoleLevel,
	roleRepository RoleRepository,
) (User, error) {
	if name == "" {
		return User{}, ErrNameRequired
	}

	role, err := roleRepository.GetByLevel(roleLevel)
	if err != nil {
		return User{}, err
	}

	return User{
		name: name,
		role: role,
	}, nil
}

type User struct {
	id      ID
	ownerID ID
	name    string
	role    Role
}

func (u User) NewChildUser(
	name string,
	roleLevel auth.RoleLevel,
	roleRepository RoleRepository,
) (User, error) {
	u, err := newUser(name, roleLevel, roleRepository)

	if err != nil {
		return User{}, err
	}

	u.ownerID = u.id

	return u, nil
}

func (u User) ID() ID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Role() Role {
	return u.role
}

func (u *User) Identify(id ID) {
	u.id = id
}

func (u *User) Rename(name string) {
	u.name = name
}

func (u User) IsOwnedBy(other User) bool {
	return u.ownerID == other.id
}

func (u User) Is(other User) bool {
	return u.id == other.id
}

type ID = int

type Repository interface {
	Add(User) (ID, error)
	Update(User) error
	Delete(ID) error
	Get(ID) (User, error)
	GetAdmin() (User, error)
	List() ([]User, error)
}
