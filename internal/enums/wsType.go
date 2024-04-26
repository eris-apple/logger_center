package enums

type WSStatus int

const (
	LogCreate WSStatus = iota
	LogSend
	SignIn
	Authorized
	Unverified
)

func (r WSStatus) String() string {
	return [...]string{"log_create", "log_send", "sign_in", "authorized", "unverified"}[r]
}
