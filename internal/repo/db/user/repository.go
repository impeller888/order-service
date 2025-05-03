package user

import (
	"context"
	"fmt"
	"local/order-service/internal/entity"
	"local/order-service/internal/repo/db"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, firstname, lastname, fullname, age, is_married, passwordhash, created_at, updated_at FROM users WHERE id = $1", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbItems, err := pgx.CollectRows(rows, pgx.RowToStructByName[user])
	if err != nil {
		return nil, err
	}

	if len(dbItems) == 0 {
		return nil, entity.ErrNotFound
	}

	return dbItems[0].toEntity(), nil
}

func (r *UserRepository) CreateUser(ctx context.Context, req *entity.UserRequest) (uuid.UUID, error) {
	var id uuid.UUID

	if req.Firstname == nil || req.Lastname == nil || req.Fullname == nil || req.Age == nil || req.IsMarried == nil {
		return uuid.Nil, db.ErrMissedRequiredField
	}

	var (
		namedArgs = pgx.NamedArgs{
			"passwordhash": req.PasswordHash,
			"firstname":    *req.Firstname,
			"lastname":     *req.Lastname,
			"fullname":     *req.Fullname,
			"age":          *req.Age,
			"is_married":   *req.IsMarried,
		}
		colNames, argNames []string
	)

	for col, _ := range namedArgs {
		colNames = append(colNames, col)
		argNames = append(argNames, "@"+col)
	}

	err := r.db.QueryRow(ctx, fmt.Sprintf("INSERT INTO users (%s) VALUES (%s) RETURNING id", strings.Join(colNames, ","), strings.Join(argNames, ",")), namedArgs).
		Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userID uuid.UUID, params *entity.UserRequest) error {
	var (
		setArgs   []string
		namedArgs = pgx.NamedArgs{
			"id": userID,
		}
	)

	if params.Firstname != nil {
		setArgs = append(setArgs, "first_name = @first_name")
		namedArgs["first_name"] = *params.Firstname
	}
	if params.Lastname != nil {
		setArgs = append(setArgs, "last_name = @last_name")
		namedArgs["last_name"] = *params.Lastname
	}

	if params.Fullname != nil {
		setArgs = append(setArgs, "full_name = @full_name")
		namedArgs["full_name"] = *params.Fullname
	}

	if params.Age != nil {
		setArgs = append(setArgs, "age = @age")
		namedArgs["age"] = *params.Age
	}

	if params.IsMarried != nil {
		setArgs = append(setArgs, "is_married = @is_married")
		namedArgs["is_married"] = *params.IsMarried
	}

	if params.PasswordHash != nil {
		setArgs = append(setArgs, "password = @password")
		namedArgs["password"] = params.PasswordHash
	}

	if len(setArgs) == 0 {
		return nil
	}

	_, err := r.db.Exec(ctx, fmt.Sprintf("UPDATE users SET %s WHERE id = @id", strings.Join(setArgs, ",")), namedArgs)

	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
