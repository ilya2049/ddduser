package user_test

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

func TestAdminCreator_CreateAdmin(t *testing.T) {
	roleRepository := memory.RoleRepository{}
	_, _ = roleRepository.Add(role.New(dict.RoleAdmin, auth.RoleLevelAdmin))

	userRepository := memory.UserRepository{}

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userJasonCredentials, err := user.NewCredentials(dict.NameJason, dict.EmailJason)
	require.NoError(t, err)

	userJason, err := adminCreator.CreateAdmin(userJasonCredentials)
	require.NoError(t, err)
	assert.Equal(t, auth.RoleLevelAdmin, userJason.Role().Level())

	userLeslieCredentials, err := user.NewCredentials(dict.NameLeslie, dict.EmailLaslie)
	require.NoError(t, err)

	_, err = adminCreator.CreateAdmin(userLeslieCredentials)
	require.Error(t, err)
	assert.EqualError(t, user.ErrOnlyOneAdmin, err.Error())
}
