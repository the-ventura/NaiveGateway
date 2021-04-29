import React from 'react'
import axios from 'axios'
import { eventBus } from '../eventBus'

export const Transfers = () => {
  return (
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
          <TableContents />
        </tbody>
      </table>
    </div>
  )
}

class TableContents extends React.Component {
  constructor (props) {
    super(props)
    this.state = { transfers: [] }
  }

  componentDidMount () {
    this.getData()
    eventBus.on('update_transfer_data', (data) =>
      this.getData()
    )
  }

  componentWillUnmount () {
    eventBus.remove('update_transfer_data')
  }

  getData () {
    axios.get(`${process.env.REACT_APP_API_URL}/v1/transactions`, {})
      .then((res) => {
        this.setState({ transfers: res.data })
      })
  }

  render () {
    return (
    // Render the state. Initially it will be null but after api result, it will display the content.
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
