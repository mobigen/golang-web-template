package models

// Todo ....
type Todo struct {
	ID      int
	Title   string
	Content string
	Status  string
}

// Todos is slice of Todo
type Todos []Todo
