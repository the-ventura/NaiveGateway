import React from 'react'
import ax from '../../requests'
import { Formik, Form, Field } from 'formik'

export class TopUp extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      account_id: '',
      amount: 0
    }
  }

  depositToAccount (id, amount) {
    ax.post('/v1/accounts/deposit', JSON.stringify({
      account_id: id,
      amount: amount
    }, null, 2)
    ).then((res) => {
      this.setState({
        account_id: id,
        amount: amount
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
            id: '',
            amount: 0
          }}
          onSubmit={async (values) => {
            await new Promise((resolve) => setTimeout(resolve, 500))
            this.depositToAccount(values.id, values.amount)
          }}
        >
          <Form className='transfer-container'>
            <div className='transfer-input'>
              <label htmlFor='id'>Account ID</label>
              <Field id='id' name='id' placeholder='Account ID' />
            </div>
            <div className='transfer-input'>
              <label htmlFor='amount'>Amount</label>
              <Field id='amount' name='amount' placeholder='Amount' />
            </div>
            <button type='submit'>Submit</button>
          </Form>
        </Formik>
        <div className='input-notification'>
          {this.state.account_id ? `Deposited ${this.state.amount} to account ${this.state.account_id}` : null}
        </div>
      </div>
    )
  }
}
