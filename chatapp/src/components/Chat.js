import React, {useState, useEffect, useRef} from 'react';
import md5 from 'md5'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {Container, Row, Col, InputGroup, InputGroupText, Input, InputGroupAddon, Button,} from 'reactstrap'
import MessagesList from './MessagesList'

var ws = null

if (window.location.host === '127.0.0.1:8000') {
  ws = new WebSocket(`ws://${window.location.host}/ws`)
} else {
  ws = new WebSocket(`ws://127.0.0.1:8000/ws`)
}

const Chat = () => {
  const [joined, setJoined] = useState(false)
  const [message, setMessage] = useState('')
  const [userName, setUserName] = useState(null)
  const [messagesList, setMsgsList] = useState([])

  const typingRef = useRef(null)

  const prevMessageList = useRef()

  const send = () => {
    setMessage('')
    const timestamp = new Date().toLocaleTimeString()
    if (message !== '') {
      const newMsg = {
        id: md5(`${timestamp}${userName}`),
        sender: {
          userName,
        },
        message,
        timestamp,
        type: "broadcast"
      }
      ws.send(JSON.stringify(newMsg))
      // chatRef.current.innerHTML += `<Row><Col md="12"><Card body style={{height: "30px"}}><CardTitle>You</CardTitle><CardSubtitle>${timestamp}</CardSubtitle><CardText>${newMsg.message}</CardText></Card></Col></Row>`
      console.log(messagesList);
      setMsgsList([...messagesList, newMsg])
      console.log(messagesList);
    }
  }

  const typing = () => {
    const newMsg = {
      sender: {
        userName
      },
      type: "typing"
    }
    ws.send(JSON.stringify(newMsg))
  }

  const join = () => {
    if (!userName) {
      toast("Enter a username")
      return
    }
    setJoined(true)
    const client = {
      userName
    }
    ws.send(JSON.stringify(client))
  }

  const handleReceivedMessage = (e) => {
    let msg = JSON.parse(e.data)
    console.log(msg);
    switch (msg.type) {
      case "join":
        return toast(`${msg.sender.userName} has joined to general chat.`)
        break;
      case "leave":
        return toast(`${msg.sender.userName} has left.`)
        break
      case "broadcast":
        console.log(messagesList);
        setMsgsList([...prevMessageList.current, msg])
        console.log(messagesList);
        break
      case "typing":
        typingRef.current.innerHTML = `${msg.sender.userName} is typing...`
        setTimeout(() => {
          typingRef.current.innerHTML = ''
        }, 2000)
        break
      default:
        console.log("no type");
    }
  }

  useEffect(() => {
    ws.addEventListener('message', handleReceivedMessage)
  }, [])

  useEffect(() => {
    prevMessageList.current = messagesList
  }, [messagesList])


  return (
    <Container fluid={true}>
     <ToastContainer/>
     <MessagesList messagesList={messagesList}/>
      <span>
        <p ref={typingRef}></p>
      </span>
      {
        joined ?
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
        :
          <Row>
            <Col md={5}>
              <InputGroup>
                <InputGroupAddon addonType="prepend">
                  <InputGroupText>username</InputGroupText>
                </InputGroupAddon>
                <Input type="text" value={userName} onChange={(e) => setUserName(e.target.value)} placeholder="username" />
              </InputGroup>
            </Col>
            <Col md={2}>
              <Button color="secondary" size="md" active onClick={() => join()}>Join</Button>
            </Col>
          </Row>
      }
    </Container>
  )
}

export default Chat
