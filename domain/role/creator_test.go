package role_test

import (
	"ddduser/db/memory"
	"ddduser/dict"
	"ddduser/domain/auth"
	"ddduser/domain/role"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreator_CreateRole(t *testing.T) {
	roleRepository := memory.RoleRepository{}
	factory := role.NewCreator(&roleRepository)

	roleAdmin := role.New(dict.RoleAdmin, auth.RoleLevelAdmin)
	roleModerator := role.New(dict.RoleModerator, auth.RoleLevelModerator)
	roleGuest := role.New(dict.RoleGuest, auth.RoleLevelGuest)

	err := factory.CreateRole(roleAdmin)
	require.NoError(t, err)

	err = factory.CreateRole(roleAdmin)
	require.Error(t, err)
	assert.EqualError(t, role.ErrRoleWithSameLevelAlreadyExists, err.Error())

	err = factory.CreateRole(roleModerator)
	require.NoError(t, err)

	err = factory.CreateRole(roleModerator)
	require.Error(t, err)
	assert.EqualError(t, role.ErrRoleWithSameLevelAlreadyExists, err.Error())

	err = factory.CreateRole(roleGuest)
	require.NoError(t, err)

	err = factory.CreateRole(roleGuest)
	require.Error(t, err)
	assert.EqualError(t, role.ErrRoleWithSameLevelAlreadyExists, err.Error())
}
