package core

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/edlingao/hexago/internal/users/ports/driven"
)

type User struct {
	ID        string    `json:"-" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserService struct {
	DBService driven.StoringUsers[User]
}

func NewUserService(db driven.StoringUsers[User]) UserService {
	return UserService{
		DBService: db,
	}
}

func (us *UserService) SignIn(username string, password string) (User, error) {
	user, err := us.DBService.GetByField("username", username, "users")

	if err != nil {
		return User{}, err
	}

	if !us.ValidatePassword(user.Password, password) {
		return User{}, errors.New("Invalid password")
	}

	return user, nil
}

func (us *UserService) Register(username, password string) error {
	if username == "" || password == "" {
		return errors.New("Username and password are required")
	}

	user := User{
		Username: username,
		Password: us.EncryptPassword(password),
	}

	query := `INSERT INTO users (username, password) VALUES (:username, :password) RETURNING id`
	err := us.DBService.Insert(user, query)
	return err
}

func (us *UserService) Get(id string) (User, error) {
	if id == "" {
		return User{}, errors.New("ID is required")
	}

	return us.DBService.Get(id, "users")
}

func (us *UserService) GetByUsername(username string) ( User, error ) {
  return us.DBService.GetByField("username", username, "users")
}

func (us *UserService) EncryptPassword(password string) string {
	plain_password := []byte(password)
	hash := sha256.Sum256(plain_password)

	return hex.EncodeToString(hash[:])
}

func (us *UserService) ValidatePassword(hash string, password string) bool {
	return us.EncryptPassword(password) == hash
}
