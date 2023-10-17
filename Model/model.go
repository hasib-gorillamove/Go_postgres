package model

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Deleted_at  int    `json:"deleted_at"`
}
