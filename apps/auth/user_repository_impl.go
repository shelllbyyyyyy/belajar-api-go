package auth

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}



func (r repository) Register(ctx context.Context, user *User) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	query := `
	INSERT INTO public.users (id, username, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	var id string
	err = tx.QueryRowContext(ctx, query,
		user.Id,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&id)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r repository) Update(ctx context.Context, payload *updateUserPayload) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}

	query := `
	UPDATE public.users
	SET username = $1,
		email = $2,
		password = $3
	WHERE id = $4`

	_, err = tx.ExecContext(ctx, query, 
		payload.Username,
		payload.Email,
		payload.Password,
		payload.Id,
	)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
	SELECT id, username, email, password, created_at, updated_at
	FROM public.users
	WHERE email = $1`

	user := &User{}
	result, err := r.db.QueryContext(ctx, query, email)

	if err != nil {
		return nil, err
	}

	result.Next()

	err = result.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r repository) FindById(ctx context.Context, id string) (*User, error) {
	query := `
	SELECT id, username, email, password, created_at, updated_at
	FROM public.users
	WHERE email = $1`

	user := &User{}
	result, err := r.db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}

	result.Next()

	err = result.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
