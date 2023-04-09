package repositories

import (
	"context"
	"database/sql"

	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/exception"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/helpers"
	"github.com/SyaibanAhmadRamadhan/impl-repo-pattern-go/models/entities"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entities.Users) entities.Users
	Update(ctx context.Context, tx *sql.Tx, user entities.Users) entities.Users
	Delete(ctx context.Context, tx *sql.Tx, userId string)
	FindById(ctx context.Context, tx *sql.Tx, userId string) entities.Users
	FindAll(ctx context.Context, tx *sql.Tx) []entities.Users
}

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user entities.Users) entities.Users {
	SQL := "SELECT id FROM users WHERE email=? LIMIT 1"
	rows, _ := tx.QueryContext(ctx, SQL, user.Email)
	if rows.Next() {
		panic(exception.NewBadRequestError("email is alvailable"))
	}
	SQL = "INSERT INTO users(id, username, email, password, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	helpers.PanicIfErr(err)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entities.Users) entities.Users {
	SQL := "UPDATE users SET username=?,updated_at=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.UpdatedAt, user.Id)
	helpers.PanicIfErr(err)
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId string) {
	SQL := "DELETE FROM users WHERE id=?"
	rows, err := tx.ExecContext(ctx, SQL, userId)
	helpers.PanicIfErr(err)
	result, err := rows.RowsAffected()
	helpers.PanicIfErr(err)
	if result == 0 {
		panic(exception.NewNotFoundError("user not found"))
	}
}

func (repostory *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) entities.Users {
	SQL := "SELECT id, username, email, created_at, updated_at FROM users WHERE id=? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helpers.PanicIfErr(err)
	defer rows.Close()

	user := entities.Users{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		helpers.PanicIfErr(err)
		return user
	} else {
		panic(exception.NewNotFoundError("user not found"))
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entities.Users {
	SQL := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfErr(err)
	defer rows.Close()

	var users []entities.Users
	for rows.Next() {
		user := entities.Users{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		helpers.PanicIfErr(err)
		users = append(users, user)
	}
	return users
}
