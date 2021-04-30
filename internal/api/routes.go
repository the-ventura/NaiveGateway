package api

import (
	"encoding/json"
	"naivegateway/internal/account"
	"naivegateway/internal/transaction"
	"net/http"
)

// Routing methods for the api

// Simple health endpoint - always returns 200 - OK
func (api *API) health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// AccountRequestBody is a generic representation of acceptable payloads
// for account operations
type AccountRequestBody struct {
	ID     string  `json:"account_id"`
	Amount float64 `json:"amount,string"`
	Name   string  `json:"account_name"`
}

// Create a new account
// Expects a payload like
// {
//		"account_name": "some_name"
// }
func (api *API) createAccount(w http.ResponseWriter, req *http.Request) {
	bodyInfo := AccountRequestBody{}
	err := json.NewDecoder(req.Body).Decode(&bodyInfo)

	account, err := account.CreateNewAccount(bodyInfo.Name, api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	payload, _ := json.Marshal(account)
	w.Write(payload)
}

// Deposit funds to an account
// Expects a payload like
// {
//		"account_id": "id",
//		"amount": 0
// }
func (api *API) depositToAccount(w http.ResponseWriter, req *http.Request) {
	bodyInfo := AccountRequestBody{}
	err := json.NewDecoder(req.Body).Decode(&bodyInfo)
	account, err := account.GetAccountByID(bodyInfo.ID, api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	account.Deposit(bodyInfo.Amount, api.db)
	payload, _ := json.Marshal(account)
	w.Write(payload)
}

// Fetches account details
// Expects a payload like
// {
//		"account_id": "id",
// }
func (api *API) accountDetails(w http.ResponseWriter, req *http.Request) {
	bodyInfo := AccountRequestBody{}
	err := json.NewDecoder(req.Body).Decode(&bodyInfo)
	account, err := account.GetAccountByID(bodyInfo.ID, api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	payload, _ := json.Marshal(account)
	w.Write(payload)
}

// AcountStatement is a payload for sending details regarding the movements in an account
type AccountStatement struct {
	AccountID            string                    `json:"account_id"`
	InboundTransactions  []transaction.Transaction `json:"inbound_transactions"`
	OutboundTransactions []transaction.Transaction `json:"outbound_transactions"`
}

// Fetches account statement
// Expects a payload like
// {
//		"account_id": "id",
// }
func (api *API) accountStatement(w http.ResponseWriter, req *http.Request) {
	bodyInfo := AccountRequestBody{}
	err := json.NewDecoder(req.Body).Decode(&bodyInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	inboundTransactions, err := transaction.ListInboundTransactionsForAccount(bodyInfo.ID, api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	outboundTransactions, err := transaction.ListOutboundTransactionsForAccount(bodyInfo.ID, api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	payload, _ := json.Marshal(AccountStatement{
		AccountID:            bodyInfo.ID,
		InboundTransactions:  inboundTransactions,
		OutboundTransactions: outboundTransactions,
	})
	w.Write(payload)
}

// TransactionRequestBody is a generic representation of acceptable payloads
// for transaction operations
type TransactionRequestBody struct {
	FromID      string  `json:"from_id"`
	ToID        string  `json:"to_id"`
	Amount      float64 `json:"amount,string"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	ID          string  `json:"transaction_id"`
}

// Creates a new transaction and places it in the ledger
// Expects a payload like
// {
//		"from_id": "id",
//		"to_id": "id",
//		"amount": "0",
//		"description": "some description",
//		"currency": "EUR",
// }
func (api *API) createTransaction(w http.ResponseWriter, req *http.Request) {
	bodyInfo := TransactionRequestBody{}
	err := json.NewDecoder(req.Body).Decode(&bodyInfo)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	opts := transaction.NewOpts()
	opts.Description = bodyInfo.Description
	if bodyInfo.Currency != "" {
		opts.Currency = bodyInfo.Currency
	}
	t := transaction.New(bodyInfo.FromID, bodyInfo.ToID, bodyInfo.Amount, opts)
	t.Stage(api.db)
	payload, _ := json.Marshal(t)
	w.Write(payload)
}

// Executes a transaction in the ledger
// Expects a payload like
// {
//		"transaction_id": "id",
// }
func (api *API) executeTransaction(w http.ResponseWriter, req *http.Request) {
	bodyInfo := TransactionRequestBody{}
	err := json.NewDecoder(req.Body).Decode(&bodyInfo)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t, err := transaction.GetTransactionByID(bodyInfo.ID, api.db)
	if err != nil {
		log.Error("Transaction not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = t.Execute(api.db)
	if err != nil {
		log.Error("Transaction failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payload, _ := json.Marshal(t)
	w.Write(payload)
}

// Gets all transaction in the ledger
func (api *API) getTransactions(w http.ResponseWriter, req *http.Request) {
	transactions, err := transaction.GetAllTransactions(api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payload, _ := json.Marshal(transactions)
	w.Write(payload)
}
