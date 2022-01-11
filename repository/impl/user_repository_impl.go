package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/repository/repository"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {

}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "SELECT * FROM user WHERE username = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Username, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()

	userData := domain.User{}
	if rows.Next() {
		err := rows.Scan(&userData.Id, &userData.FullName, &userData.Username, &userData.Password, &userData.Email, &userData.UserImg, &userData.PhoneNumber, &userData.RoleId, &userData.UserStatus)
		helper.PanicIfError(err)
		return userData, nil
	} else {
		return userData, errors.New("Akun Anda Tidak Ditemukan")
	}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO user(fullname, username, password, email, user_img, phone_number, role_id, user_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.FullName, user.Username, user.Password, user.Email, user.UserImg, user.PhoneNumber, user.RoleId, user.UserStatus)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) GetUser(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT * FROM user WHERE id =  ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	userData := domain.User{}
	if rows.Next() {
		err := rows.Scan(&userData.Id, &userData.FullName, &userData.Username, &userData.Password, &userData.Email, &userData.UserImg, &userData.PhoneNumber, &userData.RoleId, &userData.UserStatus)
		helper.PanicIfError(err)
		return userData, nil
	} else {
		return userData, errors.New("Akun Anda Tidak Ditemukan")
	}
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE user SET fullname = ?, password = ?, phone_number = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.FullName, user.Password, user.PhoneNumber, user.Id)
	helper.PanicIfError(err)

	return user
}