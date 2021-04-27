package api

import (
	"encoding/json"
	"naivegateway/internal/account"
	"net/http"
)

func (api *API) health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (api *API) createAccount(w http.ResponseWriter, req *http.Request) {
	account, err := account.CreateNewAccount(api.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	payload, _ := json.Marshal(account)
	w.Write(payload)
}

type AccountRequestBody struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount,string"`
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
