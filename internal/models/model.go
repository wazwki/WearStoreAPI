package models

type Item struct {
	Price       float64           `json:"price"`
	Title       string            `json:"title"`
	Photo       string            `json:"photo"`
	Description map[string]string `json:"description"`
}
