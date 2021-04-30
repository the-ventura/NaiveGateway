package transaction

import (
	"errors"
	"naivegateway/internal/account"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

// Transaction represents a transfer of funds between two entities
type Transaction struct {
	ID           string    `json:"id"`
	FromID       string    `json:"from_id"`
	ToID         string    `json:"to_id"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	Currency     string    `json:"currency"`
	CreationTime time.Time `json:"creation_time"`
	UUID         string    `json:"uuid" pg:",pk"`
}

// TransactionOptions is an optional data structure
type TransactionOptions struct {
	Description string
	Currency    string
}

// NewOpts is the TransactionOptions constructor
func NewOpts() TransactionOptions {
	return TransactionOptions{
		Description: "",
		Currency:    "EUR",
	}
}

// New creates a new transaction with sane defaults
func New(from, to string, amount float64, opts TransactionOptions) Transaction {
	return Transaction{
		FromID:       from,
		ToID:         to,
		Amount:       amount,
		Status:       "pending",
		Description:  opts.Description,
		Currency:     opts.Currency,
		CreationTime: time.Now(),
		UUID:         uuid.New().String(),
	}
}

// Stage places a transaction in the ledger
func (t *Transaction) Stage(db *pg.DB) error {
	// Ensure the sender exists
	sender, err := account.GetAccountByID(t.FromID, db)
	if err != nil {
		log.Error(err)
		return err
	}
	// Block funds
	err = sender.Reserve(t.Amount, db)
	if err != nil {
		log.Error(err)
		t.Delete(db)
		return err
	}
	// Create the transaction
	_, err = db.Model(t).Insert()
	if err != nil {
		log.Error(err)
	}
	return err
}

// Delete cancels a transaction
func (t *Transaction) Delete(db *pg.DB) error {
	t.Status = "canceled"
	_, err := db.Model(t).WherePK().Update()
	if err != nil {
		log.Error(err)
	}
	return nil
}

// Execute takes a transaction from the ledger and moves the listed funds
func (t *Transaction) Execute(db *pg.DB) error {
	// We need to ensure that both the sender and recipient exist
	sender, err1 := account.GetAccountByID(t.FromID, db)
	recipient, err2 := account.GetAccountByID(t.ToID, db)
	if err1 != nil || err2 != nil {
		t.Delete(db)
		return errors.New("Failed to validate transaction actors")
	}
	// Withdraw funds from the sender
	err := sender.Withdraw(t.Amount, db)
	if err != nil {
		t.Delete(db)
		return errors.New("Could not withdraw funds")
	}
	// Release blocked funds
	err = sender.Release(t.Amount, db)
	if err != nil {
		t.Delete(db)
		sender.Deposit(t.Amount, db) // Rollback withdrawn funds
		return errors.New("Could not release funds")
	}
	// Send funds to the target
	err = recipient.Deposit(t.Amount, db)
	if err != nil {
		t.Delete(db)
		sender.Deposit(t.Amount, db) // Rollback withdrawn funds
		return errors.New("Could not deposit funds")
	}
	// Mark transaction as complete
	t.Status = "completed"
	_, err = db.Model(t).WherePK().Update()
	if err != nil {
		log.Error(err)
	}
	return err
}
