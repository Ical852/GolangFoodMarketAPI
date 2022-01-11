package repository

import (
	"Go-FoodMarket/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	GetUser(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}