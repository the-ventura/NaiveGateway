package api

import (
	"encoding/json"
	"naivegateway/internal/account"
	"naivegateway/internal/transaction"
	"net/http"
)

func (api *API) health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Account methods
type AccountRequestBody struct {
	ID     string  `json:"account_id"`
	Amount float64 `json:"amount,string"`
	Name   string  `json:"account_name"`
}

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

type AccountStatement struct {
	AccountID            string                    `json:"account_id"`
	InboundTransactions  []transaction.Transaction `json:"inbound_transactions"`
	OutboundTransactions []transaction.Transaction `json:"outbound_transactions"`
}

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

// Transaction routes
type TransactionRequestBody struct {
	FromID      string  `json:"from_id"`
	ToID        string  `json:"to_id"`
	Amount      float64 `json:"amount,string"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	ID          string  `json:"transaction_id"`
}

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

func (api *API) getTransactions(w http.ResponseWriter, req *http.Request) {
	transactions, err := transaction.GetAllTransactions(api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payload, _ := json.Marshal(transactions)
	w.Write(payload)
}
