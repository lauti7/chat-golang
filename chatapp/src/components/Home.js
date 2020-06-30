import React from 'react'
import {Container, Row, Col} from 'reactstrap'
import Users from './Users'

const Home = () => {

  return (
    <Container fluid>
      <Row>
        <Col md={2}>
          <Users/>
        </Col>
      </Row>
    </Container>
  )
}

export default Home
