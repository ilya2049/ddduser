package user_test

import (
	"ddduser/db/memory"
	"ddduser/domain/auth"
	"ddduser/domain/role"
	"ddduser/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdminCreator_NewAdmin(t *testing.T) {
	roleRepository := memory.RoleRepository{}
	_, _ = roleRepository.Add(role.New("Admin", auth.RoleLevelAdmin))

	userRepository := memory.UserRepository{}

	adminCreator := user.NewAdminCreator(&roleRepository, &userRepository)

	userJason, err := adminCreator.NewAdmin("Jason")
	require.NoError(t, err)
	assert.Equal(t, auth.RoleLevelAdmin, userJason.Role().Level())

	_, err = adminCreator.NewAdmin("Leslie")
	require.Error(t, err)
	assert.EqualError(t, user.ErrOnlyOneAdmin, err.Error())
}