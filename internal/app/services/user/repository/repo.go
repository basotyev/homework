package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"lesson13/internal/app/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUserById(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int) (*models.User, error)
	RemoveUserById(ctx context.Context, id int) error
}

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) Repository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (p *PostgresUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	row := p.db.QueryRow(ctx, `INSERT INTO users(name, email, age) VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Email, user.Age)
	err := row.Scan(&user.Id)
	return err
}

func (p *PostgresUserRepository) UpdateUserById(ctx context.Context, user *models.User) error {
	_, err := p.db.Exec(ctx, `UPDATE users SET name = $1, email = $2, age = $3 WHERE id = $4`, user.Name, user.Email, user.Age, user.Id)
	return err
}

func (p *PostgresUserRepository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	var usr models.User
	row := p.db.QueryRow(ctx, `SELECT name, email, age, created_at, updated_at FROM users WHERE id = $1`, id)
	err := row.Scan(&usr.Name, &usr.Email, &usr.Age, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return nil, err
	}
	usr.Id = id
	return &usr, nil
}

func (p *PostgresUserRepository) RemoveUserById(ctx context.Context, id int) error {
	_, err := p.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}
