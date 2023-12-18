package dto

type CreateUser struct {
	Name  string
	Email string
}

type UserFilter struct {
	ID    int
	Name  string
	Email string
}
