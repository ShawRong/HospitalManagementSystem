package model

type User struct {
	Username string `json:"username"`
	IsMember bool   `json:"isMember"`
	ID       int    `json:"id"`
	Honesty  int    `json:"honesty"`
}
