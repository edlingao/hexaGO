package ports

import "github.com/edlingao/hexago/internal/users/core"

type StoringUsers interface {
	Close()
	Insert(item core.User) error
	Get(id, table string) (core.User, error)
	GetByField(field, value, table string) (core.User, error)
	DeleteByField(field, value, table string) error
	GetAll(table string) []core.User
	Delete(id, table string) error
}

type StoringSessions[T any] interface {
	Close()
	Insert(item T, sql string) error
	Get(id, table string) (T, error)
	GetByField(field, value, table string) (T, error)
	DeleteByField(field, value, table string) error
	GetAll(table string) []T
	Delete(id, table string) error
}
