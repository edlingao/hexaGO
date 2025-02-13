package driving

import (
	"github.com/edlingao/hexago/internal/users/core"
	"github.com/edlingao/hexago/internal/users/ports/driven"
)

type UserService interface {
  NewUserService (db driven.StoringUsers[core.User]) *UserService
  Register(user core.User) error
  Get(id string) (core.User, error)
  EncryptPassword(password string) string
  ValidatePassword(hash string, password string) bool
}

