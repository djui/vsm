package vsm

// Role defines a permission role.
type Role int

// String returns a human readable role text.
//
// String implements the Stringer interface.
func (r Role) String() string {
	return roleNames[r]
}

// IsMember returns if a given group of roles contains the role.
func (r Role) IsMember(group []Role) bool {
	for _, m := range group {
		if r == m {
			return true
		}
	}

	return false
}

// Enumeration of permission roles.
const (
	RoleAutomatic Role = iota
	RoleAdmin
	RoleEndUser
	RoleHunter
)

var roleNames = map[Role]string{
	RoleAutomatic: "Automatic",
	RoleAdmin:     "Admin",
	RoleEndUser:   "End-user",
	RoleHunter:    "Hunter",
}

// RoleAllUsers is a group containing all role members.
var RoleAllUsers = []Role{RoleAdmin, RoleEndUser, RoleHunter}
