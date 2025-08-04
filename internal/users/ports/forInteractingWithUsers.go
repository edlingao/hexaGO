package ports

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/edlingao/hexago/internal/users/core"
)

type UserServiceMethods interface {
	Register(username, password string) error
	SignIn(username string, password string) (core.User, error)
	Get(id string) (core.User, error)
	GetByUsername(username string) (core.User, error)
	EncryptPassword(password string) string
	ValidatePassword(hash string, password string) bool
}

type UserService struct {
	DBService StoringUsers
}

func NewUserService(db StoringUsers) UserService {
	return UserService{
		DBService: db,
	}
}

func (us UserService) SignIn(username string, password string) (core.User, error) {
	user, err := us.DBService.GetByField("username", username, "users")

	if err != nil {
		return core.User{}, err
	}

	if !us.ValidatePassword(user.Password, password) {
		return core.User{}, errors.New("Invalid password")
	}

	return user, nil
}

func (us UserService) Register(username, password string) error {
	if username == "" || password == "" {
		return errors.New("Username and password are required")
	}

	user := core.User{
		Username: username,
		Password: us.EncryptPassword(password),
	}

	err := us.DBService.Insert(user)
	return err
}

func (us UserService) Get(id string) (core.User, error) {
	if id == "" {
		return core.User{}, errors.New("ID is required")
	}

	return us.DBService.Get(id, "users")
}

func (us UserService) GetByUsername(username string) (core.User, error) {
	return us.DBService.GetByField("username", username, "users")
}

func (us UserService) EncryptPassword(password string) string {
	plain_password := []byte(password)
	hash := sha256.Sum256(plain_password)

	return hex.EncodeToString(hash[:])
}

func (us UserService) ValidatePassword(hash string, password string) bool {
	return us.EncryptPassword(password) == hash
}
