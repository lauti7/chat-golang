import React, {useReducer} from 'react';
import {AuthContext} from './auth.js'
import './App.css';
import NavBar from './components/NavBar'
import Home from './components/Home'
import Login from './components/Login'


const authId = JSON.parse(localStorage.getItem("authId"))

const initialState = {
  isAuthenticated: authId ? true : false,
  authId: authId ? authId : null
}

const reducer = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      localStorage.setItem("authId", JSON.stringify(action.authId))
      return {
        isAuthenticated: true,
        authId: action.authId
      }
    case "LOGOUT":
      localStorage.clear()
      return {
        isAuthenticated: false,
        authId: null
      }
  }
}


const App = () => {

  const [state, dispatch] = useReducer(reducer, initialState)

  return (
    <AuthContext.Provider value={{state, dispatch}} >
      <div className="App">
        <NavBar/>
        {
          state.isAuthenticated ?
            <Home/>
          : <Login/>
        }
      </div>
    </AuthContext.Provider>
  );
}

export default App;
