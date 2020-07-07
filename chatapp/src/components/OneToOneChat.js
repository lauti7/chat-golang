import React, {useState, useEffect} from 'react'
import {useAuth} from '../auth'
import {Container, Row, Col, InputGroup, InputGroupText, Input, InputGroupAddon, Button,} from 'reactstrap'


const OneToOneChat = () => {

  const {state, dispatch} = useAuth()

  const [message, setMessage] = useState('')
  const [chat, setChat] = useState({})


  return (
    <Container>
     <div>
      <p>User: </p>
     </div>
     <MessagesList messagesList={messagesList}/>
      <span>
        <p ref={typingRef}></p>
      </span>
      <Row>
        <Col md={8}>
          <InputGroup>
            <InputGroupAddon addonType="prepend">
              <InputGroupText>New Message</InputGroupText>
            </InputGroupAddon>
            <Input value={message} onChange={(e) => {typing();setMessage(e.target.value)}} />
          </InputGroup>
        </Col>
        <Col md={4}>
          <Button color="secondary" size="lg" active onClick={() => send()}>Send</Button>
        </Col>
      </Row>
    </Container>
  )


}
