package enums

type Role int

const (
	Admin Role = iota
	Moderator
	User
	Guest
)

func (r Role) String() string {
	return [...]string{"admin", "moderator", "user", "guest"}[r]
}
