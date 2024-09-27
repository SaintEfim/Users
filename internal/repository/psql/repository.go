package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"Users/config"
	"Users/internal/models/entity"
	"Users/internal/models/interfaces"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Repository struct {
	db  *sql.DB
	cfg *config.Config
}

func NewRepository(db *sql.DB, cfg *config.Config) interfaces.Repository {
	return &Repository{
		db:  db,
		cfg: cfg,
	}
}

func (r *Repository) Get(ctx context.Context) ([]*entity.UserEntity, error) {
	var users []*entity.UserEntity

	rows, err := r.db.Query(retrieveAllUsers)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		user := &entity.UserEntity{}
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, fmt.Errorf("row scan error: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}

func (r *Repository) GetOneById(ctx context.Context, id string) (*entity.UserEntity, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, fmt.Errorf("invalid UUID: %v", err)
	}

	user := &entity.UserEntity{}

	if err := r.db.QueryRow(retrieveOneById, id).Scan(&user.Id, &user.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user found with id: %s", id)
		}
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}

	return user, nil
}

func (r *Repository) Create(ctx context.Context, user *entity.UserEntity) error {
	user.Id, _ = uuid.NewRandom()
	if _, err := r.db.Exec(createUser, user.Name); err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid UUID: %v", err)
	}

	result, err := r.db.Exec(deleteUser, id)
	if err != nil {
		return fmt.Errorf("error executing delete query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error retrieving rows affected count: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id: %s", id)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, id string, user *entity.UserEntity) error {

	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid UUID: %v", err)
	}

	result, err := r.db.Exec(updateUser, user.Name, id)
	if err != nil {
		return fmt.Errorf("error executing update query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error retrieving rows affected count: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id: %s", id)
	}

	return nil
}
