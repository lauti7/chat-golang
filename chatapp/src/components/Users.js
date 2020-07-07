import React, {useState, useEffect} from 'react'
import {useAuth} from '../auth'
import {Container, Row, Col} from 'reactstrap'
import UsersList  from './UsersList'


const Users = () => {

  const {state, dispatch} = useAuth()

  const [users, setUsers] = useState([])

  const fetchUsers = () => {
    fetch("http://localhost:8080/api/users", {
      method: "GET",
      headers: {
        "Authorization": state.authId
      }
    })
    .then(res => res.json())
    .then(json => {
      setUsers([...json.users])
    })
  }

  useEffect(() => {
    fetchUsers()
  }, [])

  return (
    <div className="mt-3">
      <UsersList users={users}/>
    </div>
  )
}

export default Users
