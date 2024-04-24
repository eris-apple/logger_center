package store

type Store interface {
	User() UserRepository
	Session() SessionRepository
	Project() ProjectRepository
	Log() LogRepository
	ServiceAccount() ServiceAccountRepository
}
