package store

type Store interface {
	User() UserRepository
	Project() ProjectRepository
	Session() SessionRepository
}
