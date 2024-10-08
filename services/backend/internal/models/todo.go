package models

type TodoList struct {
	ID 			int 	`json:"id"`
	Title 		string	`json:"title"`
	Description	string	`json:"description"`
}

type UsersList struct {
	ID		int
	UserID	int
	ListID	int
}

type TodoItem struct {
	ID 			int		`json:"id"`
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	Done		bool	`json:"done"`
}

type ListItem struct {
	ID		int
	UserID	int
	ListID	int
}