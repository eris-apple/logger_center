package enums

type UserStatus int

const (
	Pending UserStatus = iota
	Accepted
	Declined
	Banned
)

func (r UserStatus) String() string {
	return [...]string{"pending", "accepted", "declined", "banned"}[r]
}
