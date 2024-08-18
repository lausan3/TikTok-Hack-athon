package entities

type Notification[T any] struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Type     string `json:"type"`
	Content  T      `json:"content"`
}
