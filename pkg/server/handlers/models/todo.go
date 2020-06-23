package models

// Todo keeps fields for request/response handling.
type Todo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	IsFinished bool   `json:"is_finished"`
}
