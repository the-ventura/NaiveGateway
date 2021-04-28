import axios from 'axios'
import { Formik, Form, Field } from 'formik'
import NavBar from '../../components/NavBar'
import './Home.css'

const Home = () => {
  return (
    <div className='main'>
      <NavBar />
      <div>
        <h3>New transfer</h3>
        <div>
          <Formik
            initialValues={{
              sender: '',
              recipient: '',
              amount: ''
            }}
            onSubmit={async (values) => {
              await new Promise((resolve) => setTimeout(resolve, 500))
              axios.post(`${process.env.API_URL}/v1/transactions/create`, JSON.stringify(values, null, 2)).then((response) => {
                console.log(response)
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
        <h3>Pending transfers</h3>
        <div className='transfers'>
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
              <tr className='transfer-row'>
                <th className='text'>142390db-8389-41da-bb6f-8f814b76cf8a</th>
                <th className='text'>aead8895-13a4-41f7-8f1d-a2dafa2618f8</th>
                <th className='text'>0289e84b-e0b7-480e-ba28-91fad415b59f</th>
                <th className='numeric'>1000</th>
                <th className='text'>Pending</th>
              </tr>
              <tr className='transfer-row'>
                <th className='text'>142390db-8389-41da-bb6f-8f814b76cf8a</th>
                <th className='text'>aead8895-13a4-41f7-8f1d-a2dafa2618f8</th>
                <th className='text'>0289e84b-e0b7-480e-ba28-91fad415b59f</th>
                <th className='numeric'>1000</th>
                <th className='text'>Pending</th>
              </tr>
            </tbody>
          </table>
        </div>
        <h3>Approve transfers</h3>
        <div>
          <Formik
            initialValues={{
              transferId: ''
            }}
            onSubmit={async (values) => {
              await new Promise((resolve) => setTimeout(resolve, 500))
              axios.post(`${process.env.API_URL}/v1/transactions/execute`, JSON.stringify(values, null, 2)).then((response) => {
                console.log(response)
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
      </div>
    </div>
  )
}

export default Home
