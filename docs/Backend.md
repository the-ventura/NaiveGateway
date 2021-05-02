# Backend

## Design

The backend service contains an http api so that other services can communicate with it.
It is also in charge of managing the state of transactions and accounts.

State is stored in a postgres database since it is performant, reliable and its relational nature reflects the strucutre of the data pretty well

## Technology

The entire service is written in golang, cobra is used to parse command line arguments and go-pg is used to communicate with the database.
For convenience gorilla/mux is used to serve http content.

Everything else is homemade :)

## API

### /health

Description:
Health endpoint to check if the service is running

Method: GET

Response: 200 OK

### /v1/accounts/create

Description:
Creates a new account

Method: POST

Expected Payload:

```json
{
  "account_name": "some_name"
}
```

Response:

```json
{
  "id": "",
  "uuid": "",
  "available": "",
  "blocked": "",
  "deposited": "",
  "withdrawn": "",
  "currency": "",
  "card_name": "",
  "card_type": "",
  "card_number": "",
  "card_expiry_month": "",
  "card_expiry_year": "",
  "card_security_code": "",
  "creation_time": ""
}
```

### /v1/accounts/deposit

Description:
Deposit funds to an account

Method: POST

Expected Payload:

```json
{
  "account_id": "id",
  "amount": 0
}
```

Response:

```json
{
  "id": "",
  "uuid": "",
  "available": "",
  "blocked": "",
  "deposited": "",
  "withdrawn": "",
  "currency": "",
  "card_name": "",
  "card_type": "",
  "card_number": "",
  "card_expiry_month": "",
  "card_expiry_year": "",
  "card_security_code": "",
  "creation_time": ""
}
```

### /v1/accounts/detail

Description:
Get account details

Method: POST

Expected Payload:

```json
{
  "account_id": "id",
}
```

Response:

```json
{
  "id": "",
  "uuid": "",
  "available": "",
  "blocked": "",
  "deposited": "",
  "withdrawn": "",
  "currency": "",
  "card_name": "",
  "card_type": "",
  "card_number": "",
  "card_expiry_month": "",
  "card_expiry_year": "",
  "card_security_code": "",
  "creation_time": ""
}
```

### /v1/accounts/statement

Description:
Get account statement

Method: POST

Expected Payload:

```json
{
  "account_id": "id",
}

```json
Response:

```json
{
  "account_id": "",
  "inbound_transactions": [],
  "outbound_transactions": []
}

}
```

### /v1/transactions

Description:
Gets all transactions

Method: GET

Response:

```json
{
  [
    "id": "",
    "from_id": "",
    "to_id": "",
    "amount": "",
    "status": "",
    "description": "",
    "currency": "",
    "creation_time": "",
    "uuid": ""
  ]
}
```

### /v1/transactions/create

Description:
Creates a new transaction

Method: POST

Expected Payload:

```json
{
  "from_id": "id",
  "to_id": "id",
  "amount": "0",
  "description": "some description",
  "currency": "EUR",
}
```

Response:

```json
{
  "id": "",
  "from_id": "",
  "to_id": "",
  "amount": "",
  "status": "",
  "description": "",
  "currency": "",
  "creation_time": "",
  "uuid": ""
}
```

### /v1/transactions/execute

Description:
Executes a pending transaction

Method: POST

Expected Payload:

```json
{
  "transaction_": "id",
}
```

Response:

```json
{
  "id": "",
  "from_id": "",
  "to_id": "",
  "amount": "",
  "status": "",
  "description": "",
  "currency": "",
  "creation_time": "",
  "uuid": ""
}
```
