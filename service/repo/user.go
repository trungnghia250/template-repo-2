package repo

import (
	"SecondAssignment/service/model"
	"context"
	"database/sql"
	"fmt"
)

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *sql.DB
}

type UserRepo interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	GetAll(ctx context.Context) ([]model.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user *model.User) error
}

func (u *userRepo) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	row := u.db.QueryRowContext(ctx,
		`SELECT id, username, email, phone, date_of_birth
			   FROM users
			   WHERE id = $1`, id)

	err := row.Scan(&user.ID, &user.UserName, &user.Phone, &user.Email, &user.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("error get user id: %w", err)
	}

	return &user, nil
}

func (u *userRepo) GetAll(ctx context.Context) ([]model.User, error) {
	rows, err := u.db.QueryContext(ctx, `SELECT id, username, email, phone, date_of_birth 
											   FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.UserName, &user.Phone, &user.Email, &user.DateOfBirth)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	query := "UPDATE users SET username = ?, email = ?, phone = ?, date_of_birth = ? WHERE id = ?"
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error update user id: %w", err)
	}
	_, err = stmt.ExecContext(ctx, user.UserName, user.Email, user.Phone, user.DateOfBirth, user.ID)
	if err != nil {
		return fmt.Errorf("error update user id: %w", err)
	}

	return nil
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	_, err := u.db.ExecContext(
		ctx,
		`DELETE FROM users where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete user from the database: %w", err)
	}

	return nil
}

func (u *userRepo) Create(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (id, username, email, phone, date_of_birth) VALUES (?, ?, ?, ?, ?)"
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error create user id: %w", err)
	}
	_, err = stmt.ExecContext(ctx, user.ID, user.UserName, user.Email, user.Phone, user.DateOfBirth)
	if err != nil {
		return fmt.Errorf("error create user id: %w", err)
	}

	return nil
}
