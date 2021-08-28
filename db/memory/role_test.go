package memory_test

import (
	"ddduser/db/memory"
	"ddduser/dict"
	"ddduser/domain/auth"
	"ddduser/domain/role"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleRepository_GetByLevel(t *testing.T) {
	repository := memory.RoleRepository{}

	_, _ = repository.Add(role.New(dict.RoleModerator, auth.RoleLevelModerator))
	_, _ = repository.Add(role.New(dict.RoleGuest, auth.RoleLevelGuest))

	moderator, _ := repository.GetByLevel(auth.RoleLevelModerator)
	assert.Equal(t, auth.RoleLevelModerator, moderator.Level())

	_, err := repository.GetByLevel(auth.RoleLevelAdmin)
	assert.EqualError(t, role.ErrRoleDoesNotExist, err.Error())
}
