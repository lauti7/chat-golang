import React from 'react'
import { ListGroup, ListGroupItem, Badge, Button } from 'reactstrap';

const UsersList = ({users}) => {
  return (
    <ListGroup>
      {
        users.map(user => {
          return (
            <ListGroupItem>
              <div className="d-flex justify-content-between">
                <p className="m-0 p-0">{user.user_name}</p>
                <Button size="sm">Select</Button>
              </div>
            </ListGroupItem>
          )
        })
      }
    </ListGroup>
  )
}

export default UsersList
