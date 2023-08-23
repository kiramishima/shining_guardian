package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	dbErrors "github.com/kiramishima/shining_guardian/pkg/errors"
	"github.com/kiramishima/shining_guardian/src/domain"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(conn *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: conn,
	}
}

// FindByCredentials repository method for retrieve user if provide credentials are valid
func (repo *AuthRepository) FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.User, error) {
	var query = `SELECT id,
		   email,
		   password,
		   (SELECT id
			FROM roles
					 INNER JOIN public.model_has_roles mhr on roles.id = mhr.role_id
			WHERE mhr.model_id = u.id
			LIMIT 1) role_id,
		   created_at,
		   updated_at
	FROM users u
	WHERE email = $1 AND password = $2`
	stmt, err := repo.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", dbErrors.ErrPrepareStatement, err)
	}
	defer stmt.Close()

	u := &domain.User{}

	row := stmt.QueryRowContext(ctx, data.Email)
	var updatedAt sql.NullTime
	err = row.Scan(&u.ID, &u.Email, &u.Password, &u.RoleID, &u.CreatedAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, dbErrors.ErrUserNotFound
		} else {
			return nil, fmt.Errorf("%s: %w", dbErrors.ErrScanData, err)
		}
	}

	if updatedAt.Valid {
		u.UpdatedAt = updatedAt.Time
	}

	return u, nil
}

// Register repository method for create a new user.
func (repo *AuthRepository) Register(ctx context.Context, registerReq *domain.RegisterRequest) error {
	var query = `INSERT INTO users(email, password) VALUES($1, $2)`
	stmt, err := repo.db.PreparexContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", dbErrors.ErrPrepareStatement, err)
	}
	defer stmt.Close()

	registerReq.Password = registerReq.Hash256Password(registerReq.Password)
	_, err = stmt.QueryxContext(ctx, registerReq.Email, registerReq.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dbErrors.ErrAlreadyExists
		} else {
			return fmt.Errorf("%s: %w", dbErrors.ErrScanData, err)
		}
	}

	return nil
}
