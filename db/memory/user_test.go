package memory_test

import (
	"ddduser/db/memory"
	"ddduser/dict"
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

	roleCreator := role.NewCreator(&roleRepository)
	err := roleCreator.CreateTestRoles()
	require.NoError(t, err)

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userAdmin, err := adminCreator.CreateAdmin(dict.NameLeslie)
	require.NoError(t, err)

	userSam, err := userAdmin.NewChildUser(dict.NameSam, auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userJames, err := userSam.NewChildUser(dict.NameJames, auth.RoleLevelGuest, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(userSam)
	userJamesID, _ := userRepository.Add(userJames)

	users, _ := userRepository.List()

	const (
		usersInRepository = 3
	)

	var (
		userAdminIdx = userAdmin.ID() - 1
		userSamIdx   = userSamID - 1
		userJamesIdx = userJamesID - 1
	)

	require.Len(t, users, usersInRepository)
	assert.Equal(t, users[userAdminIdx].Name(), userAdmin.Name())
	assert.Equal(t, users[userSamIdx].Name(), userSam.Name())
	assert.Equal(t, users[userJamesIdx].Name(), userJames.Name())
}

func TestUserRepository_Delete(t *testing.T) {
	userRepository := memory.UserRepository{}
	roleRepository := memory.RoleRepository{}

	roleCreator := role.NewCreator(&roleRepository)
	err := roleCreator.CreateTestRoles()
	require.NoError(t, err)

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userAdmin, err := adminCreator.CreateAdmin(dict.NameLeslie)
	require.NoError(t, err)

	userSam, err := userAdmin.NewChildUser(dict.NameSam, auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userJames, err := userSam.NewChildUser(dict.NameJames, auth.RoleLevelGuest, &roleRepository)
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

	roleCreator := role.NewCreator(&roleRepository)
	err := roleCreator.CreateTestRoles()
	require.NoError(t, err)

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userAdmin, err := adminCreator.CreateAdmin(dict.NameLeslie)
	require.NoError(t, err)

	userSam, err := userAdmin.NewChildUser(dict.NameSam, auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(userSam)

	userSam, err = userRepository.Get(userSamID)
	require.NoError(t, err)

	userSam.Rename(dict.NameSamantha)

	err = userRepository.Update(userSam)
	require.NoError(t, err)

	userSam, err = userRepository.Get(userSamID)
	require.NoError(t, err)

	assert.Equal(t, dict.NameSamantha, userSam.Name())
}
