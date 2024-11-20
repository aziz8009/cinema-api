package repositories

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/aziz8009/cinema-app/src/entities"
)

type UserRepository interface {
	GetAll(ctx context.Context) (res []entities.User, err error)
	GetByEmail(ctx context.Context, email string) (res entities.User, err error)
	Create(ctx context.Context, req entities.User) (res entities.User, err error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {

	if db == nil {
		panic("db is nil")
	}

	fmt.Println("masuk sini 1")

	return &userRepository{db: db}
}

func (r *userRepository) GetAll(ctx context.Context) (res []entities.User, err error) {
	query, args, err := sq.Select("id", "name", "email").
		From("users").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return res, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (res entities.User, err error) {

	query, args, err := sq.Select("id", "name", "email", "password").
		From("users").
		Where("email = ?", email).
		ToSql()

	if err != nil {
		return res, err
	}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(&res.ID, &res.Name, &res.Email, &res.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case where no rows are found
			return res, nil
		}
		return res, err
	}

	// Return the result
	return res, nil
}

func (r *userRepository) Create(ctx context.Context, user entities.User) (res entities.User, err error) {
	// Prepare the SQL query to insert the user
	query, args, err := sq.Insert("users").
		Columns("name", "email", "password").
		Values(user.Name, user.Email, user.Password).
		ToSql()

	if err != nil {
		return res, err
	}

	// Execute the query
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return res, err
	}

	// Get the last inserted ID (assuming your `id` column is auto-incrementing)
	userID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res = user
	res.ID = int64(userID)

	return res, nil
}
