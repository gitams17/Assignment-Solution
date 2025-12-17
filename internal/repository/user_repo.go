package repository

import (
	"context"
	db "user-api/db/sqlc" // Import the generated code
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, id int32) (db.User, error)
	ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type SQLStore struct {
	*db.Queries
}

func NewRepository(q *db.Queries) UserRepository {
	return &SQLStore{Queries: q}
}