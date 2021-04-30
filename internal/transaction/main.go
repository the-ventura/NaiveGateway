package transaction

import (
	"naivegateway/internal/logger"

	"github.com/go-pg/pg/v10"
)

var log = logger.Log

// Generic transaction operations

// GetTransactionByID gets a transaction object by its id
func GetTransactionByID(id string, db *pg.DB) (*Transaction, error) {
	t := Transaction{}
	t.UUID = id
	err := db.Model(&t).Where("uuid = ?", id).Select()
	if err != nil {
		log.Error(err)
	}
	return &t, err
}

// GetAllTransactions fetches all transactions
func GetAllTransactions(db *pg.DB) ([]Transaction, error) {
	var transactions []Transaction
	err := db.Model(&transactions).Order("creation_time ASC").Select()
	if err != nil {
		log.Error(err)
	}
	return transactions, err
}

// ListInboundTransactionsForAccount gets all inbound transactions for a given account
func ListInboundTransactionsForAccount(id string, db *pg.DB) ([]Transaction, error) {
	var inboundTransactions []Transaction
	err := db.Model(&inboundTransactions).Where("to_id = ?", id).Order("creation_time ASC").Select()
	if err != nil {
		log.Error(err)
	}
	return inboundTransactions, err
}

// ListOutboundTransactionsForAccount gets all outbound transactions for a given account
func ListOutboundTransactionsForAccount(id string, db *pg.DB) ([]Transaction, error) {
	var outboundTransactions []Transaction
	err := db.Model(&outboundTransactions).Where("from_id = ?", id).Order("creation_time ASC").Select()
	if err != nil {
		log.Error(err)
	}
	return outboundTransactions, err
}
