package enum

// UserStatus should map to user_status_enum
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

// AllCategories returns all possible values for Category
func (u UserStatus) All() []string {
	return []string{
		string(UserStatusActive),
		string(UserStatusInactive),
	}
}
