package account

import (
	"naivegateway/internal/logger"

	"github.com/go-pg/pg/v10"
)

var log = logger.Log

func CreateNewAccount(db *pg.DB) (*Account, error) {
	account := New()
	err := account.Create(db)
	return &account, err
}

func GetAccountByID(id string, db *pg.DB) (*Account, error) {
	a := Account{}
	a.UUID = id
	err := db.Model(&a).Where("uuid = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return &a, err
}
