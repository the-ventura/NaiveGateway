import React, { useState } from 'react'
import { Route, Switch, BrowserRouter } from 'react-router-dom'
import axios from 'axios'
import Home from './pages/Home'
import NavBar from './components/NavBar'
const App = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path='/accounts'>
          <Ping />
        </Route>
        <Route exact path='/'>
          <Home />
        </Route>
      </Switch>
    </BrowserRouter>
  )
}

const Ping = () => {
  const [notification, setNotification] = useState('')

  const handlePing = async () => {
    try {
      const response = await axios.get('/api/ping')
      setNotification(`Successful ping with response: ${response.data}`)
    } catch (e) {
      setNotification('Failed to ping')
    }
    setTimeout(() => setNotification(''), 2000)
  }

  return (
    <div>
      <NavBar />
      <div>
        <p>{notification}</p>
        <button onClick={handlePing}>Ping</button>
      </div>
    </div>
  )
}

export default App
