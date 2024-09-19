package user

import (
	"acme/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	DB *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}
func (repo *PostgresUserRepository) GetUsers() ([]model.User, error) {
	users := []model.User{}
	err := sqlx.Select(repo.DB, &users, "SELECT * FROM users")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.User{}, errors.New("database could not be queried")
	}
	return users, nil
}

func (repo *PostgresUserRepository) AddUser(user model.User) (id int, err error) {
	err = repo.DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", user.Name).Scan(&id)
	if err != nil {
		fmt.Println("Error inserting user into the database:", err)
		return 0, errors.New("could not insert user")
	}
	return id, nil
}

// GetUser retrieves a user by ID from the database.
func (repo *PostgresUserRepository) GetUser(id int) (model.User, error) {
	// Implement logic to fetch a user by ID from the database
	var user model.User
	err := sqlx.Get(repo.DB, &user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error querying the database", err)
		return model.User{}, errors.New("database could not be queried")
	}
	return user, nil
}
func (repo *PostgresUserRepository) UpdateUser(id int, user *model.User) (model.User, error) {
	// Implement logic to update a user in the database
	return model.User{}, nil
}

// DeleteUser deletes a user from the database.
func (repo *PostgresUserRepository) DeleteUser(id int) error {
	// Implement logic to delete a user from the database
	_, err := repo.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error deleting user from the database:", err)
		return errors.New("could not delete user")
	}
	return nil
}

func (repo *PostgresUserRepository) Close() {
}
