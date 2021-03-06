import ax from '../../requests'
import { Formik, Form, Field } from 'formik'
import { eventBus } from '../eventBus'

export const ApproveTransfers = () => {
  return (
    <div>
      <Formik
        initialValues={{
          transferId: ''
        }}
        onSubmit={async (values) => {
          await new Promise((resolve) => setTimeout(resolve, 500))
          ax.post('/v1/transactions/execute', JSON.stringify({
            transaction_id: values.transferId
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
            <label htmlFor='transferId'>Transfer ID</label>
            <Field id='transferId' name='transferId' placeholder='Transfer ID' />
          </div>
          <button type='submit'>Submit</button>
        </Form>
      </Formik>
    </div>
  )
}
