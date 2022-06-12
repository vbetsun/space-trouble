package psql

import (
	"database/sql"
	"fmt"

	"github.com/vbetsun/space-trouble/internal/core"
)

// User represents repository for users entity
type User struct {
	db *sql.DB
}

// NewUser return instance of user repository
func NewUser(db *sql.DB) *User {
	return &User{db}
}

// CreateUser creates new user in DB
func (r *User) CreateUser(u core.User) (core.User, error) {
	var user core.User
	err := r.db.QueryRow(createUserQuery(), u.FirstName, u.LastName, u.Gender, u.Birthday).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Gender, &user.Birthday)
	return user, err
}

// GetUser returns user from DB by name
func (r *User) GetUser(firstName, lastName string) (core.User, error) {
	var user core.User
	err := r.db.QueryRow(getUserQuery(), firstName, lastName).Scan(&user.ID)

	return user, err
}

func createUserQuery() string {
	return fmt.Sprintf(`--sql
		INSERT INTO %s (first_name, last_name, gender, birthday) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, first_name, last_name, gender, birthday
	`, usersTable)
}

func getUserQuery() string {
	return fmt.Sprintf(`--sql
		SELECT id FROM %s 
		WHERE first_name = $1 
		AND last_name = $2
	`, usersTable)
}
