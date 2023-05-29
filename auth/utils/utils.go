package utils

type User struct {
	Name     string `json:"full_name,omitempty" form:"full_name"`
	Username string `json:"username,omitempty" form:"username"`
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	Action   string `json:"action,omitempty" form:"action"`
}
