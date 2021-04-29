import NavBar from '../../components/NavBar'
import { NewAccount, AccountDetails, TopUp } from '../../components/Accounts'

const Accounts = () => {
  return (
    <div className='main'>
      <NavBar />
      <div>
        <h3>Create</h3>
        <NewAccount />
        <h3>Details</h3>
        <AccountDetails />
        <h3>Top Up</h3>
        <TopUp />
      </div>
    </div>
  )
}

export default Accounts
