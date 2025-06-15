package adapters

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SessionsStore[Item any] struct {
	db *sqlx.DB
}

func NewSessionStore[Item any]() *SessionsStore[Item] {
	db, err := sqlx.Connect("sqlite3", "./db/main.db")
	if err != nil {
		log.Fatal(err)
	}

	return &SessionsStore[Item]{db: db}
}

func (s *SessionsStore[Item]) Close() {
	s.db.Close()
}

func (s *SessionsStore[Item]) Insert(item Item, sql string) error {
	_, err := s.db.NamedExec(sql, item)
	return err
}

func (s *SessionsStore[Item]) Get(id, table string) (Item, error) {
	var item Item
	err := s.db.Get(&item, "SELECT * FROM "+table+" WHERE id = ?", id)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *SessionsStore[Item]) GetAll(table string) []Item {
	var items []Item
	err := s.db.Select(&items, "SELECT * FROM "+table+" ORDER BY id DESC")

	if err != nil {
		log.Fatal(err)
	}

	return items
}

func (s *SessionsStore[Item]) Delete(id, table string) error {
	_, err := s.db.Exec("DELETE FROM "+table+" WHERE id = ?", id)
	return err
}

func (s *SessionsStore[Item]) DeleteByField(field, value, table string) error {
	_, err := s.db.Exec("DELETE FROM "+table+" WHERE "+field+" = ?", value)

	return err
}

func (s *SessionsStore[Item]) GetByField(field, value, table string) (Item, error) {
	var item Item
	err := s.db.Get(&item, "SELECT * FROM "+table+" WHERE "+field+" = ?", value)

	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *SessionsStore[Item]) GetSQL(query string, item Item) (Item, error) {
	err := s.db.Get(&item, query)
	if err != nil {
		return item, err
	}
	return item, nil
}
