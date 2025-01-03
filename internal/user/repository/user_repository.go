package repository

import (
	"GolangProject/internal/user"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FetchAllUsers() ([]user.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM tbl_users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
