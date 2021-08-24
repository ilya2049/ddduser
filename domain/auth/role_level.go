package auth

type RoleLevel string

const (
	RoleLevelAdmin     RoleLevel = "admin"
	RoleLevelModerator RoleLevel = "moderator"
	RoleLevelGuest     RoleLevel = "guest"
)
