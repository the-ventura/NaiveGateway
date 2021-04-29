import axios from 'axios'
import { Formik, Form, Field } from 'formik'
import { eventBus } from '../eventBus'

export const NewTransfer = () => {
  return (
    <div>
      <Formik
        initialValues={{
          sender: '',
          recipient: '',
          amount: ''
        }}
        onSubmit={async (values) => {
          await new Promise((resolve) => setTimeout(resolve, 500))
          axios.post(`${process.env.REACT_APP_API_URL}/v1/transactions/create`, JSON.stringify({
            from_id: values.sender,
            to_id: values.recipient,
            amount: values.amount
          }, null, 2)
          ).then(() => {
            eventBus.dispatch('update_transfer_data', {})
          }, (error) => {
            console.log(error)
          })
        }}
      >
        <Form className='transfer-container'>
          <div className='transfer-input'>
            <label htmlFor='sender'>Sender</label>
            <Field id='sender' name='sender' placeholder='Sender ID' />
          </div>
          <div className='transfer-input'>
            <label htmlFor='recipient'>Recipient</label>
            <Field id='recipient' name='recipient' placeholder='Recipient ID' />
          </div>
          <div className='transfer-input'>
            <label htmlFor='amount'>Amount</label>
            <Field id='amount' name='amount' placeholder='100' />
          </div>
          <button type='submit'>Submit</button>
        </Form>
      </Formik>
    </div>
  )
}
