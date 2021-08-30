package role_test

import (
	"context"
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

	ctx := context.Background()

	err := factory.CreateRole(ctx, roleAdmin)
	require.NoError(t, err)

	err = factory.CreateRole(ctx, roleAdmin)
	require.Error(t, err)
	assert.EqualError(t, role.ErrRoleWithSameLevelAlreadyExists, err.Error())

	err = factory.CreateRole(ctx, roleModerator)
	require.NoError(t, err)

	err = factory.CreateRole(ctx, roleModerator)
	require.Error(t, err)
	assert.EqualError(t, role.ErrRoleWithSameLevelAlreadyExists, err.Error())

	err = factory.CreateRole(ctx, roleGuest)
	require.NoError(t, err)

	err = factory.CreateRole(ctx, roleGuest)
	require.Error(t, err)
	assert.EqualError(t, role.ErrRoleWithSameLevelAlreadyExists, err.Error())
}
