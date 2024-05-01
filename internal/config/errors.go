package config

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Default
var (
	ErrBadRequest = &APIError{
		Code:    "BAD_REQUEST",
		Message: "bad request",
	}
	ErrForbiddenAccess = &APIError{
		Code:    "FORBIDDEN_ACCESS",
		Message: "forbidden access",
	}
	ErrInternalServerError = &APIError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: "internal server error",
	}
)

// Users
var (
	ErrUserNotFound = &APIError{
		Code:    "USER_NOT_FOUND",
		Message: "user not found",
	}
	ErrUsersNotFound = &APIError{
		Code:    "USERS_NOT_FOUND",
		Message: "users not found",
	}
	ErrUserAlreadyExist = &APIError{
		Code:    "USER_ALREADY_EXIST",
		Message: "user with this email already exist",
	}
	ErrIncorrectEmailOrPassword = &APIError{
		Code:    "INCORRECT_EMAIL_OR_PASSWORD",
		Message: "incorrect email or password",
	}
	ErrUserNotModerated = &APIError{
		Code:    "USER_NOT_MODERATED",
		Message: "user not moderated",
	}
	ErrUserDeclined = &APIError{
		Code:    "USER_DECLINED",
		Message: "user is declined",
	}
	ErrUserBanned = &APIError{
		Code:    "USER_BANNED",
		Message: "user is banned",
	}
	ErrUserSearchParam = &APIError{
		Code:    "USER_SEARCH_PARAM",
		Message: "search param must be a non-empty string and length greater than 3",
	}
)

// Sessions
var (
	ErrSessionExpired = &APIError{
		Code:    "SESSION_EXPIRED",
		Message: "session expired",
	}
)

// Projects
var (
	ErrProjectNotFound = &APIError{
		Code:    "PROJECT_NOT_FOUND",
		Message: "project not found",
	}
	ErrProjectsNotFound = &APIError{
		Code:    "PROJECTS_NOT_FOUND",
		Message: "projects not found",
	}
)

// Logs
var (
	ErrLogNotFound = &APIError{
		Code:    "LOG_NOT_FOUND",
		Message: "log not found",
	}
	ErrLogsNotFound = &APIError{
		Code:    "LOGS_NOT_FOUND",
		Message: "logs not found",
	}
)

// ServiceAccount
var (
	ErrServiceAccountNotFound = &APIError{
		Code:    "SERVICE_ACCOUNT_NOT_FOUND",
		Message: "service account not found",
	}
	ErrServiceAccountsNotFound = &APIError{
		Code:    "SERVICE_ACCOUNTS_NOT_FOUND",
		Message: "service accounts not found",
	}
)
