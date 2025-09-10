package services

import (
	"database/sql"
	"errors"
	"log"
	"warehouse/backend/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(user *models.User) error {
	var storedPassword string
	err := s.db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		log.Println(err)
		return errors.New("internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}

func (s *AuthService) Register(user *models.User) error {
	// Basic validation
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password are required")
	}

	// Check if user already exists
	var existingUser string
	err := s.db.QueryRow("SELECT username FROM users WHERE username = ?", user.Username).Scan(&existingUser)
	if err != sql.ErrNoRows {
		return errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	// Store new user
	stmt, err := s.db.Prepare("INSERT INTO users(username, password) VALUES(?,?)")
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	_, err = stmt.Exec(user.Username, string(hashedPassword))
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	return nil
}