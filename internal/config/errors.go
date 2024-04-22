package config

// Default
var (
	ErrBadRequest          = "bad request"
	ErrForbiddenAccess     = "forbidden access"
	ErrInternalServerError = "internal server error"
)

// Users
var (
	ErrUserNotFound             = "user not found"
	ErrUsersNotFound            = "users not found"
	ErrUserAlreadyExist         = "user with this email already exist"
	ErrIncorrectEmailOrPassword = "incorrect email or password"
)

// Sessions
var (
	ErrSessionExpired = "session expired"
)

// Projects
var (
	ErrProjectNotFound  = "project not found"
	ErrProjectsNotFound = "projects not found"
)
