package account

import (
	"naivegateway/internal/logger"

	"github.com/go-pg/pg/v10"
)

var log = logger.Log

// Generic Account manipulation

// CreateNewAccount creates a new account with the specified name
func CreateNewAccount(name string, db *pg.DB) (*Account, error) {
	account := New()
	account.CardName = name
	err := account.Create(db)
	return &account, err
}

// GetAccountByID fetches an account object by a given ID
func GetAccountByID(id string, db *pg.DB) (*Account, error) {
	a := Account{}
	a.UUID = id
	err := db.Model(&a).Where("uuid = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return &a, err
}
