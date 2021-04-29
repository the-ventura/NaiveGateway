import NavBar from '../../components/NavBar'
import { Transfers, NewTransfer, ApproveTransfers } from '../../components/Transfers'
import './Home.css'

const Home = () => {
  return (
    <div className='main'>
      <NavBar />
      <div>
        <h3>New transfer</h3>
        <NewTransfer />
        <h3>Pending transfers</h3>
        <Transfers />
        <h3>Approve transfers</h3>
        <ApproveTransfers />
      </div>
    </div>
  )
}

export default Home
