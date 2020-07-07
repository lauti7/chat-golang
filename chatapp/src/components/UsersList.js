import React from 'react'
import {useAuth} from '../auth'
import { ListGroup, ListGroupItem, Badge, Button } from 'reactstrap';

const UsersList = ({users}) => {

  const selectChat = id => {
    console.log(id);
  }

  return (
    <ListGroup>
      {
        users.map(user => {
          return (
            <ListGroupItem>
              <div className="d-flex justify-content-between">
                <p className="m-0 p-0">{user.user_name}</p>
                <Button size="sm" onClick={() => selectChat(user.id)}>Select</Button>
              </div>
            </ListGroupItem>
          )
        })
      }
    </ListGroup>
  )
}

export default UsersList
