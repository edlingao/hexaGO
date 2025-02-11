package adapters

import (
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

type Store[Item any] struct {
  db *sqlx.DB
}

func NewDB[Item any]() *Store[Item] {
  db, err := sqlx.Connect("sqlite3", "./db/main.db")
  if err != nil {
    log.Fatal(err)
  }

  return &Store[Item]{db: db}
}

func (s *Store[Item]) Close() {
  s.db.Close()
}

func (s *Store[Item]) Insert(item Item, sql string) error {
  _, err := s.db.NamedExec(sql, item)
  return err
}

func (s *Store[Item]) Get(id, table string) ( Item, error ) {
  var item Item
  err := s.db.Get(&item, "SELECT * FROM " + table + " WHERE id = ?", id)
  if err != nil {
    return item, err
  }
  return item, nil
}

func (s *Store[Item]) GetAll(table string) []Item {
  var items []Item
  err := s.db.Select(&items, "SELECT * FROM " + table + " ORDER BY id DESC")

  if err != nil {
    log.Fatal(err)
  }

  return items
}

func (s *Store[Item]) Delete(id, table string) error {
  _, err := s.db.Exec("DELETE FROM " + table + " WHERE id = ?", id)
  return err
}
