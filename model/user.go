package model

type User struct {
	UserID   uint64 `json:"user_id,omitempty" db:"user_id"`
	UserName string `json:"username,omitempty" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
}
