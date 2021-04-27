package account

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type Account struct {
	ID               int64     `json:"id"`
	UUID             string    `json:"uuid" pg:",pk"`
	Available        float64   `json:"available"`
	Blocked          float64   `json:"blocked"`
	Deposited        float64   `json:"deposited"`
	Withdrawn        float64   `json:"withdrawn"`
	Currency         string    `json:"currency"`
	CardName         string    `json:"card_name"`
	CardType         string    `json:"card_type"`
	CardNumber       int       `json:"card_number"`
	CardExpiryMonth  int       `json:"card_expiry_month"`
	CardExpiryYear   int       `json:"card_expiry_year"`
	CardSecurityCode int       `json:"card_security_code"`
	CreationTime     time.Time `json:"creation_time"`
}

// A very simple account with a lot of assumptions.
// We could expand on this a bit but for now this model will do
func New() Account {
	var identifier = rand.Intn(4666778181156223-4666000000000000) + 4666000000000000
	return Account{
		UUID:             uuid.New().String(),
		Available:        0,
		Blocked:          0,
		Deposited:        0,
		Withdrawn:        0,
		Currency:         "EUR",
		CardName:         "Mr Payment",
		CardType:         "VISA",
		CardNumber:       identifier,
		CardExpiryMonth:  rand.Intn(12-1) + 1,
		CardExpiryYear:   rand.Intn(24-18) + 18,
		CardSecurityCode: rand.Intn(999-100) + 100,
		CreationTime:     time.Now(),
	}
}

func (a Account) String() string {
	return fmt.Sprintf("Account<%s %f %s %s>", a.UUID, a.Available, a.Currency, a.CardType)
}

func (a *Account) Create(db *pg.DB) error {
	_, err := db.Model(a).Insert()
	if err != nil {
		log.Error(err)
	}
	return err
}

func (a *Account) UpdateAvailableFunds() {
	a.Available = a.Deposited - a.Withdrawn - a.Blocked
}

func (a *Account) Deposit(amount float64, db *pg.DB) error {
	if amount < 0 {
		return errors.New("You can only deposit positive quantities")
	}
	a.Deposited += amount
	a.UpdateAvailableFunds()
	_, err := db.Model(a).WherePK().Update()
	if err != nil {
		log.Error(err)
	}
	return err
}

func (a *Account) Withdraw(amount float64, db *pg.DB) error {
	if amount < 0 {
		return errors.New("You can only withdraw positive quantities")
	}
	if amount > a.Available {
		return errors.New("You cannot withdraw more than you have available")
	}
	a.Withdrawn += amount
	a.UpdateAvailableFunds()
	_, err := db.Model(a).WherePK().Update()
	if err != nil {
		log.Error(err)
	}
	return err
}

func (a *Account) Reserve(amount float64, db *pg.DB) error {
	if amount < 0 {
		return errors.New("You can only reserve positive quantities")
	}
	if amount > a.Available {
		return errors.New("You cannot reserve more than you have available")
	}
	a.Blocked += amount
	a.UpdateAvailableFunds()
	_, err := db.Model(a).WherePK().Update()
	if err != nil {
		log.Error(err)
	}
	return err
}

func (a *Account) Release(amount float64, db *pg.DB) error {
	if amount < 0 {
		return errors.New("You can only release positive quantities")
	}
	a.Blocked -= amount
	a.UpdateAvailableFunds()
	_, err := db.Model(a).WherePK().Update()
	if err != nil {
		log.Error(err)
	}
	return err
}
