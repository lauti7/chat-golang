import React, {useReducer, useEffect} from 'react';
import {AuthContext} from './auth.js'
import './App.css';
import NavBar from './components/NavBar'
import Home from './components/Home'
import Login from './components/Login'

const connectToWS = (id) => {
  var ws = new WebSocket(`ws://127.0.0.1:8080/api/ws?user_id=${id}`)
  return ws
}

const authId = JSON.parse(localStorage.getItem("authId"))
const userName = JSON.parse(localStorage.getItem("userName"))


const initialState = {
  users: [],
  isAuthenticated: authId ? authId : false,
  authId: authId ? authId : null,
  userName: userName ? userName : null,
  currentChat: null
}

if (initialState.isAuthenticated) {
  let ws = connectToWS(initialState.authId)
  initialState.ws = ws
} else {
  initialState.ws = null
}
const reducer = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      localStorage.setItem("authId", JSON.stringify(action.authId))
      localStorage.setItem("userName", JSON.stringify(action.userName))
      let ws = connectToWS(action.authId)
      console.log(action);
      return {
        isAuthenticated: true,
        authId: action.authId,
        userName: action.userName,
        ws: ws,
        currentChat: null
      }
    case "LOGOUT":
      localStorage.clear()
      //TODO: disconnct WebSocket
      return {
        ...state,
        isAuthenticated: false,
        authId: null,
        useState: null,
        ws:null
      }
    case "SELECTCHAT":
      return {
        ...state,
        currentChat: action.currentChat
      }
    case "FETCHUSERS":
      console.log(action.users);
      return {
        ...state,
        users: [...action.users]
      }
    case "UPDATEUSERS":
      console.log(state.users);
      const idxUser = state.users.findIndex(u => u.id === action.newUser)
      console.log(idxUser)
      let users = [...state.users]
      if (action.status === 'on') {
        users[idxUser] = {...users[idxUser], online: true}
        return {
          ...state,
          users: [...users]
        }
      } else {
        users[idxUser] = {...users[idxUser], online: false}
        return {
          ...state,
          users: [...users]
        }
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
