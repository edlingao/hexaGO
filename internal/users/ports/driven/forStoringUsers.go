package driven

type StoringUsers[T any] interface {
	Close()
	Insert(item T, sql string) error
	Get(id, table string) (T, error)
  GetByField(field, value, table string) (T, error)
  DeleteByField(field, value, table string) error
	GetAll(table string) []T
  Delete(id, table string) error
}
