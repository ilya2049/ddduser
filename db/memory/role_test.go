package memory_test

import (
	"context"
	"ddduser/db/memory"
	"ddduser/dict"
	"ddduser/domain/auth"
	"ddduser/domain/role"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleRepository_GetByLevel(t *testing.T) {
	ctx := context.Background()

	repository := memory.RoleRepository{}

	_, _ = repository.Add(ctx, role.New(dict.RoleModerator, auth.RoleLevelModerator))
	_, _ = repository.Add(ctx, role.New(dict.RoleGuest, auth.RoleLevelGuest))

	moderator, _ := repository.GetByLevel(ctx, auth.RoleLevelModerator)
	assert.Equal(t, auth.RoleLevelModerator, moderator.Level())

	_, err := repository.GetByLevel(ctx, auth.RoleLevelAdmin)
	assert.EqualError(t, role.ErrRoleDoesNotExist, err.Error())
}
