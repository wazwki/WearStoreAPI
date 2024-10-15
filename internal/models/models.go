package models

type Item struct {
	Price       float64           `json:"price"`
	Title       string            `json:"title"`
	Photo       string            `json:"photo"`
	Description map[string]string `json:"description"`
}

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"fistname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}
