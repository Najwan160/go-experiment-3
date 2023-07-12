package model

type Auth struct {
	AccountID string `json:"account_id"`
	RoleID    string `json:"role_id"`
	Role      string `json:"role"`
}
