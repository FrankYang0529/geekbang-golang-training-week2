package storage

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type ItemSroarage interface {
	GetItem(itemID string) (Item, error)
}

type itemStorage struct {
	db *sql.DB
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewItemStorage(db *sql.DB) (*itemStorage, error) {
	return &itemStorage{
		db: db,
	}, nil
}

func (i *itemStorage) GetItem(itemID string) (*Item, error) {
	rows, err := i.db.Query("SELECT id, name FROM items WHERE id = ?", itemID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("storage: can't get item by id %s", itemID))
	}
	defer rows.Close()

	item := &Item{}
	err = rows.Scan(item.ID, item.Name)
	if err != nil {
		return nil, errors.Wrap(err, "storage: can't scan rows")
	}
	return item, nil
}
