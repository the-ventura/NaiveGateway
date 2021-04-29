import React from 'react'
import { Route, Switch, BrowserRouter } from 'react-router-dom'
import Home from './pages/Home'
import Accounts from './pages/Accounts'

const App = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path='/accounts'>
          <Accounts />
        </Route>
        <Route exact path='/'>
          <Home />
        </Route>
      </Switch>
    </BrowserRouter>
  )
}

export default App
