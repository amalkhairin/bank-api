package models

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

type History struct {
	UserID string  `json:"user_id"`
	Action string  `json:"action"`
	Amount float64 `json:"amount,omitempty"`
	Time   string  `json:"time"`
}
