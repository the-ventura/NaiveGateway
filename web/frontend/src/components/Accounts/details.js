import React from 'react'
import axios from 'axios'
import { Formik, Form, Field } from 'formik'

export class AccountDetails extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      show_details: false,
      details: {
        available: 0,
        blocked: 0,
        card_name: 0,
        currency: 0,
        deposited: 0,
        uuid: 0,
        withdrawn: 0
      },
      statement: {
        inbound: [],
        outbound: []
      }
    }
  }

  getAccountDetails (id) {
    axios.post(`${process.env.REACT_APP_API_URL}/v1/accounts/detail`, JSON.stringify({
      account_id: id
    }, null, 2)
    ).then((res) => {
      this.setState({
        details: {
          available: res.data.available,
          blocked: res.data.blocked,
          card_name: res.data.card_name,
          currency: res.data.currency,
          deposited: res.data.deposited,
          uuid: res.data.uuid,
          withdrawn: res.data.withdrawn
        }
      })
    }, (error) => {
      console.log(error)
    })
    axios.post(`${process.env.REACT_APP_API_URL}/v1/accounts/statement`, JSON.stringify({
      account_id: id
    }, null, 2)
    ).then((res) => {
      this.setState({
        show_details: true,
        statement: {
          inbound: res.data.inbound_transactions,
          outbound: res.data.outbound_transactions
        }
      })
    }, (error) => {
      console.log(error)
    })
  }

  render () {
    return (
      <div>
        <Formik
          initialValues={{
            account_id: ''
          }}
          onSubmit={async (values) => {
            await new Promise((resolve) => setTimeout(resolve, 500))
            this.getAccountDetails(values.account_id)
          }}
        >
          <Form className='transfer-container'>
            <div className='transfer-input'>
              <label htmlFor='account_id'>Account ID</label>
              <Field id='account_id' name='account_id' placeholder='Account ID' />
            </div>
            <button type='submit'>Submit</button>
          </Form>
        </Formik>
        <div className='input-notification'>
          {this.state.show_details ? <Balance details={this.state.details} inbound={this.state.statement.inbound} outbound={this.state.statement.outbound} /> : null}
        </div>
      </div>
    )
  }
}

const Balance = (props) => {
  console.log(props)
  return (
    <div className='transfers' style={{ height: '50vh' }}>
      <div>Account holder: {props.details.card_name}</div>
      <div>Account ID: {props.details.uuid}</div>
      <div>Available funds: {props.details.available}</div>
      <div>Unavailable funds: {props.details.blocked}</div>
      <div>Deposited funds:{props.details.deposited}</div>
      <div>Withdrawn funds:{props.details.withdrawn}</div>
      <div>Currency: {props.details.currency}</div>
      <h3> Inbound Transactions </h3>
      <Transfers transactions={props.inbound} />
      <h3> Outbound Transactions </h3>
      <Transfers transactions={props.outbound} />
    </div>
  )
}

const Transfers = (props) => {
  return (
    <table className='transfer-list'>
      <thead>
        <tr className='transfer-headers'>
          <th className='text'>Transfer ID</th>
          <th className='text'>Sender ID</th>
          <th className='text'>Recipient ID</th>
          <th className='numeric'>Amount</th>
          <th className='text'>Status</th>
        </tr>
      </thead>
      <tbody>
        <TableContents transactions={props.transactions} />
      </tbody>
    </table>
  )
}

class TableContents extends React.Component {
  constructor (props) {
    super(props)
    this.state = { transfers: props.transactions || [] }
  }

  render () {
    return (
      this.state.transfers.map((t) => {
        return (
          <tr key={t.uuid} className='transfer-row'>
            <th className='text'>{t.uuid}</th>
            <th className='text'>{t.from_id}</th>
            <th className='text'>{t.to_id}</th>
            <th className='numeric'>{t.amount}</th>
            <th className='text'>{t.status}</th>
          </tr>
        )
      })
    )
  }
}
