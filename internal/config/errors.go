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
	ErrUserNotModerated         = "user not moderated"
	ErrUserDeclined             = "user was declined"
	ErrUserBanned               = "user was banned"
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

// Logs
var (
	ErrLogNotFound  = "log not found"
	ErrLogsNotFound = "logs not found"
)

// Logs
var (
	ErrServiceAccountsNotFound = "service accounts not found"
	ErrServiceAccountNotFound  = "service account not found"
)
