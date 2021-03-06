package memory_test

import (
	"context"
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
	ctx := context.Background()

	userRepository := memory.UserRepository{}
	roleRepository := memory.RoleRepository{}

	roleCreator := role.NewCreator(&roleRepository)
	err := roleCreator.CreateTestRoles(ctx)
	require.NoError(t, err)

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userAdminCredentials, err := user.NewCredentials(dict.NameLeslie, dict.EmailLaslie)
	require.NoError(t, err)

	userAdmin, err := adminCreator.CreateAdmin(ctx, userAdminCredentials)
	require.NoError(t, err)

	userSamCredentials, err := user.NewCredentials(dict.NameSam, dict.EmailSam)
	require.NoError(t, err)

	userSam, err := userAdmin.NewChildUser(userSamCredentials, auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userJamesCredentials, err := user.NewCredentials(dict.NameJames, dict.EmailJames)
	require.NoError(t, err)

	userJames, err := userSam.NewChildUser(userJamesCredentials, auth.RoleLevelGuest, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(ctx, userSam)
	userJamesID, _ := userRepository.Add(ctx, userJames)

	users, _ := userRepository.List(ctx)

	const (
		usersInRepository = 3
	)

	var (
		userAdminIdx = userAdmin.ID() - 1
		userSamIdx   = userSamID - 1
		userJamesIdx = userJamesID - 1
	)

	require.Len(t, users, usersInRepository)
	assert.Equal(t, users[userAdminIdx].Credentials().Name(), userAdmin.Credentials().Name())
	assert.Equal(t, users[userSamIdx].Credentials().Name(), userSam.Credentials().Name())
	assert.Equal(t, users[userJamesIdx].Credentials().Name(), userJames.Credentials().Name())
}

func TestUserRepository_Delete(t *testing.T) {
	ctx := context.Background()

	userRepository := memory.UserRepository{}
	roleRepository := memory.RoleRepository{}

	roleCreator := role.NewCreator(&roleRepository)
	err := roleCreator.CreateTestRoles(ctx)
	require.NoError(t, err)

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userAdminCredentials, err := user.NewCredentials(dict.NameLeslie, dict.EmailLaslie)
	require.NoError(t, err)

	userAdmin, err := adminCreator.CreateAdmin(ctx, userAdminCredentials)
	require.NoError(t, err)

	userSamCredentials, err := user.NewCredentials(dict.NameSam, dict.EmailSam)
	require.NoError(t, err)

	userSam, err := userAdmin.NewChildUser(userSamCredentials, auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userJamesCredentials, err := user.NewCredentials(dict.NameJames, dict.EmailJames)
	require.NoError(t, err)

	userJames, err := userSam.NewChildUser(userJamesCredentials, auth.RoleLevelGuest, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(ctx, userSam)
	_, _ = userRepository.Add(ctx, userJames)

	_, err = userRepository.Get(ctx, userSamID)
	require.NoError(t, err)

	err = userRepository.Delete(ctx, userSamID)
	require.NoError(t, err)

	_, err = userRepository.Get(ctx, userSamID)
	assert.EqualError(t, user.ErrUserDoesNotExist, err.Error())
}

func TestUserRepository_Update(t *testing.T) {
	ctx := context.Background()

	userRepository := memory.UserRepository{}
	roleRepository := memory.RoleRepository{}

	roleCreator := role.NewCreator(&roleRepository)
	err := roleCreator.CreateTestRoles(ctx)
	require.NoError(t, err)

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userAdminCredentials, err := user.NewCredentials(dict.NameLeslie, dict.EmailLaslie)
	require.NoError(t, err)

	userAdmin, err := adminCreator.CreateAdmin(ctx, userAdminCredentials)
	require.NoError(t, err)

	userSamCredentials, err := user.NewCredentials(dict.NameSam, dict.EmailSam)
	require.NoError(t, err)

	userSam, err := userAdmin.NewChildUser(userSamCredentials, auth.RoleLevelModerator, &roleRepository)
	require.NoError(t, err)

	userSamID, _ := userRepository.Add(ctx, userSam)

	userSam, err = userRepository.Get(ctx, userSamID)
	require.NoError(t, err)

	userSamNewCredentials, err := user.NewCredentials(dict.NameSamantha, dict.EmailSam)
	require.NoError(t, err)

	userSam.UpdateCredentials(userSamNewCredentials)

	err = userRepository.Update(ctx, userSam)
	require.NoError(t, err)

	userSam, err = userRepository.Get(ctx, userSamID)
	require.NoError(t, err)

	assert.Equal(t, dict.NameSamantha, userSam.Credentials().Name())
}
