package store

import "github.com/kodaf1/go-basic-rest-api/internal/app/model"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.Password,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, password FROM users WHERE email=$1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Password,
	); err != nil {
		return nil, err
	}

	return u, nil
}
