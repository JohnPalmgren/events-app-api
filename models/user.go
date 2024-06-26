package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save inserts a new user record into the database.
// It hashes the user's password and stores the hashed password along with the email.
// If the insertion is successful, it sets the ID of the User object.
// It returns any error encountered during the process.
func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

// ValidateUser checks if the provided email and password match a user in the database.
// It retrieves the user's ID and hashed password based on the email,
// then compares the provided password with the stored hashed password.
// If the credentials are valid, it sets the ID of the User object.
// It returns an error if the credentials are invalid.
func (u *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CompareHashAndPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil

}
