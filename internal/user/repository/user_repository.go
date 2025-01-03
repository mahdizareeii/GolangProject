package repository

import (
	"GolangProject/internal/user"
	"database/sql"
	"fmt"
)

const tblUsers = "tbl_users"

type UserRepository struct {
	db *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() (userResult []user.User, errorResult error) {
	query := fmt.Sprintf("SELECT * FROM %s", tblUsers)
	rows, err := r.db.Query(query)
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

func (r *UserRepository) AddNewUser(u *user.User) (userIdResult int, errorResult error) {
	query := fmt.Sprintf("INSERT INTO %s (name, email) VALUES ($1,$2) RETURNING id", tblUsers)
	var id int
	err := r.db.QueryRow(query, &u.Name, &u.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
