package ports

type StoringOperations[T any] interface {
	Close()
	Insert(item T) error
	Get(id, table string) (T, error)
	GetAll(table string) []T
	Delete(id, table string) error
}
