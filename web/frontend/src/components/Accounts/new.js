import React from 'react'
import axios from 'axios'
import { Formik, Form, Field } from 'formik'

export class NewAccount extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      account_id: ''
    }
  }

  createAccount (name) {
    axios.post(`${process.env.REACT_APP_API_URL}/v1/accounts/create`, JSON.stringify({
      account_name: name
    }, null, 2)
    ).then((res) => {
      this.setState({
        account_id: res.data.uuid
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
            name: ''
          }}
          onSubmit={async (values) => {
            await new Promise((resolve) => setTimeout(resolve, 500))
            this.createAccount(values.name)
          }}
        >
          <Form className='transfer-container'>
            <div className='transfer-input'>
              <label htmlFor='name'>Name</label>
              <Field id='name' name='name' placeholder='Name' />
            </div>
            <button type='submit'>Submit</button>
          </Form>
        </Formik>
        <div className='input-notification'>
          {this.state.account_id ? `Created account ${this.state.account_id}` : null}
        </div>
      </div>
    )
  }
}
