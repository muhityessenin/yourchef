package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"user_service/internal/domain/user"
	"user_service/internal/repository/interfaces"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) Login(email, password string) (int, error) {
	var user user.Entity
	row, err := u.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return http.StatusNotFound, err
	}
	defer row.Close()
	if row.Next() {
		if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address, &user.RegistrationDate, &user.Role); err != nil {
			return http.StatusBadRequest, err
		}
	}

	if user.Id == "" {
		return http.StatusNotFound, sql.ErrNoRows
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return http.StatusUnauthorized, err
	}
	return http.StatusOK, nil
}

func (u *UserRepository) CreateUser(input *user.InputResponse) (int, error) {
	_, err := u.db.Exec("INSERT INTO users (name, email, password, address, registration_date, role) VALUES ($1, $2, $3, $4, $5, $6)",
		input.Name, input.Email, input.Password, input.Address, input.RegistrationDate, input.Role)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
}

func (u *UserRepository) UpdateUser(id string, input *user.InputResponse) (int, error) {
	_, err := u.GetUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, sql.ErrNoRows
		}
		return http.StatusInternalServerError, err
	}
	_, err = u.db.Exec("UPDATE users SET name = $1, email = $2, password = $3, address = $4, registration_date = $5, role = $6 WHERE id = $7",
		&input.Name, &input.Email, &input.Password, &input.Address, &input.RegistrationDate, &input.Role, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
}

func (u *UserRepository) DeleteUser(id string) (int, error) {
	_, err := u.GetUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, sql.ErrNoRows
		}
		return http.StatusInternalServerError, err
	}
	_, err = u.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
}

func (u *UserRepository) SearchUser(queryType, query string) ([]user.Entity, error) {
	var users []user.Entity
	var err error
	var q string
	if queryType == "name" {
		q = fmt.Sprintf("SELECT * FROM users WHERE name = $1")
	} else if queryType == "email" {
		q = fmt.Sprintf("SELECT * FROM users WHERE email = $1")
	}
	rows, err := u.db.Query(q, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var entity user.Entity
		if err := rows.Scan(&entity.Id, &entity.Name, &entity.Email, &entity.Password, &entity.Address, &entity.RegistrationDate, &entity.Role); err != nil {
			return users, err
		}
		users = append(users, entity)
	}
	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}
	return users, nil
}

func (u *UserRepository) GetUsers() ([]user.Entity, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []user.Entity
	for rows.Next() {
		var user user.Entity
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address, &user.RegistrationDate, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}
	return users, nil
}
func (u *UserRepository) GetUserById(id string) (user.Entity, error) {
	var user user.Entity
	row, err := u.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return user, err
	}
	defer row.Close()
	if row.Next() {
		if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address, &user.RegistrationDate, &user.Role); err != nil {
			return user, err
		}
	}
	if user.Id == "" {
		return user, sql.ErrNoRows
	}
	return user, nil
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{
		db: db,
	}
}
