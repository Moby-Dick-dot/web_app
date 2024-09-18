package request

type User struct {
	UserID   uint64 `json:"user_id,omitempty" db:"user_id"`
	UserName string `json:"username,omitempty" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
}

type RegisterReq struct {
	UserName        string `json:"username,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}
