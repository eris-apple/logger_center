package config

import "errors"

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Default
var (
	ErrBadRequest          = errors.New("BAD_REQUEST")
	ErrForbiddenAccess     = errors.New("FORBIDDEN_ACCESS")
	ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
)

// Users
var (
	ErrUserNotFound             = errors.New("USER_NOT_FOUND")
	ErrUsersNotFound            = errors.New("USERS_NOT_FOUND")
	ErrUserAlreadyExist         = errors.New("USER_ALREADY_EXIST")
	ErrIncorrectEmailOrPassword = errors.New("INCORRECT_EMAIL_OR_PASSWORD")
	ErrUserNotModerated         = errors.New("USER_NOT_MODERATED")
	ErrUserDeclined             = errors.New("USER_DECLINED")
	ErrUserBanned               = errors.New("USER_BANNED")
	ErrUserSearchParam          = errors.New("USER_SEARCH_PARAM")
)

// Identity
var (
	ErrInvalidEmail    = errors.New("INVALID_EMAIL")
	ErrInvalidPassword = errors.New("INVALID_PASSWORD")
)

// Sessions
var (
	ErrSessionExpired = errors.New("SESSION_EXPIRED")
)

// Projects
var (
	ErrProjectNotFound  = errors.New("PROJECT_NOT_FOUND")
	ErrProjectsNotFound = errors.New("PROJECTS_NOT_FOUND")
)

var (
	ErrInvalidProjectName   = errors.New("INVALID_PROJECT_NAME")
	ErrInvalidProjectPrefix = errors.New("INVALID_PROJECT_PREFIX")
)

// Logs
var (
	ErrLogNotFound  = errors.New("LOG_NOT_FOUND")
	ErrLogsNotFound = errors.New("LOGS_NOT_FOUND")
)

var (
	ErrInvalidLogLevel   = errors.New("INVALID_LOG_LEVEL")
	ErrInvalidLogChainID = errors.New("INVALID_LOG_CHAIN_ID")
)

// ServiceAccount
var (
	ErrServiceAccountNotFound  = errors.New("SERVICE_ACCOUNT_NOT_FOUND")
	ErrServiceAccountsNotFound = errors.New("SERVICE_ACCOUNTS_NOT_FOUND")
)

var (
	ErrInvalidServiceAccountName = errors.New("INVALID_SERVICE_ACCOUNT_NAME")
)
