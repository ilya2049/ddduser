package memory_test

import (
	"ddduser/db/memory"
	"ddduser/domain/auth"
	"ddduser/domain/role"
	"ddduser/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_List(t *testing.T) {
	userRepository := memory.UserRepository{}
	roleRepository := memory.RoleRepository{}

	_, _ = roleRepository.Add(role.New("Moderator", auth.RoleLevelModerator))
	_, _ = roleRepository.Add(role.New("Arbitrary user", auth.RoleLevelGuest))

	userSam, err := user.New("Sam", auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userJames, err := user.New("James", auth.RoleLevelGuest, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(userSam)
	userJamesID, _ := userRepository.Add(userJames)

	users, _ := userRepository.List()

	userSam.Identify(userSamID)
	userJames.Identify(userJamesID)

	assert.Equal(t, []user.User{userSam, userJames}, users)
}

func TestUserRepository_Delete(t *testing.T) {
	userRepository := memory.UserRepository{}
	roleRepository := memory.RoleRepository{}

	_, _ = roleRepository.Add(role.New("Moderator", auth.RoleLevelModerator))
	_, _ = roleRepository.Add(role.New("Arbitrary user", auth.RoleLevelGuest))

	userSam, err := user.New("Sam", auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userJames, err := user.New("James", auth.RoleLevelGuest, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(userSam)
	_, _ = userRepository.Add(userJames)

	_, err = userRepository.Get(userSamID)
	require.NoError(t, err)

	err = userRepository.Delete(userSamID)
	require.NoError(t, err)

	_, err = userRepository.Get(userSamID)
	assert.EqualError(t, user.ErrUserDoesNotExist, err.Error())
}

func TestUserRepository_Update(t *testing.T) {
	userRepository := memory.UserRepository{}
	roleRepository := memory.RoleRepository{}

	_, _ = roleRepository.Add(role.New("Moderator", auth.RoleLevelModerator))
	_, _ = roleRepository.Add(role.New("Arbitrary user", auth.RoleLevelGuest))

	userSam, err := user.New("Sam", auth.RoleLevelGuest, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(userSam)

	userSam, err = userRepository.Get(userSamID)
	require.NoError(t, err)

	assert.Equal(t, auth.RoleLevelGuest, userSam.Role().Level())

	err = userSam.ChangeRole(auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	err = userRepository.Update(userSam)
	require.NoError(t, err)

	userSam, err = userRepository.Get(userSamID)
	require.NoError(t, err)

	assert.Equal(t, auth.RoleLevelModerator, userSam.Role().Level())
}
