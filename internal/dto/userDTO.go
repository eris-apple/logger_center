package dto

type UpdateUserDTO struct {
	Email  string `json:"email"`
	Status string `json:"status"`
	Role   string `json:"role"`
}

type FindUsersDTO struct {
	Email  string `form:"email,omitempty"`
	Status string `form:"status,omitempty"`
	Role   string `form:"role,omitempty"`
}
