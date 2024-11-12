package models

import (
	"context"
	"fmt"
	"log"
	"time"

	db "github.com/bibhu20031/CollabPad/backend/db"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(username string, hashedPassword string) error {
	pool, err := db.ConnectDB()
	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}
	defer pool.Close()

	var exists bool
	err = pool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking existing user: %w", err)
	}
	if exists {
		return fmt.Errorf("username already exists")
	}

	query := `INSERT INTO users(username,password,created_at) VALUES($1,$2,$3)`
	_, err = pool.Exec(context.Background(), query, username, hashedPassword, time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func FindUserByUsername(username string) (User, error) {
	pool, err := db.ConnectDB()
	if err != nil {
		log.Println("Failed to connect to database in FindUserByUsername")
		return User{}, fmt.Errorf("database connection failed: %w", err)
	}

	var user User
	query := `SELECT id, username, password, created_at FROM users WHERE username=$1`
	err = pool.QueryRow(context.Background(), query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return User{}, fmt.Errorf("failed to find user: %w", err)
	}
	return user, nil
}
