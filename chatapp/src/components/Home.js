import React from 'react'
import {useAuth} from '../auth'
import {Container, Row, Col} from 'reactstrap'
import Users from './Users'
import OneToOneChat from './OneToOneChat'

const Home = () => {

  const {state, dispatch} = useAuth()

  return (
    <Container fluid>
      <Row>
        <Col md={2}>
          <Users/>
        </Col>
        <Col md={10}>
          {
            state.currentChat ?
              <OneToOneChat/>
            : ''
          }
        </Col>
      </Row>
    </Container>
  )
}

export default Home
