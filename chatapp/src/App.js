import React, {useState, useEffect, useRef} from 'react';
import logo from './logo.svg';
import './App.css';
import md5 from 'md5'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function App() {

  const [ws, setWs] = useState(null)
  const [joined, setJoined] = useState(false)
  const [message, setMessage] = useState('')
  const [chat, setChatMessages] = useState([])
  const [email, setEmail] = useState(null)
  const [userName, setUserName] = useState(null)
  const [status, setStatus] = useState('typing')

  const chatRef = useRef(null)
  const typingRef = useRef(null)

  const send = () => {
    if (message != '') {
      ws.send(JSON.stringify({email, userName, message, status: 'sent'}))
      chatRef.current.innerHTML += `<div className="chip">
                                      <img src="${gravatarUrl(email)}">
                                        ${userName}
                                    </div>
                                    ${message} <br/>`
     setMessage('')
    }
  }

  const typing = () => {
    ws.send(JSON.stringify({userName, email, status}))
  }



  const gravatarUrl = (email) => {
    return `https://www.gravatar.com/avatar/${md5(email)}?d=wavatar`
  }

  const join = () => {
    if (!email) {
      toast("Enter an email")
      return
    }
    if (!userName) {
      toast("Enter a username")
      return
    }
    // setEmail(email)
    // setUserName(userName)
    setJoined(true)
  }

  const handleBroadcasted = (e) => {
    let msg = JSON.parse(e.data)
    if (msg.status == "typing") {
      typingRef.current.innerText = `${msg.username} is typing...`
      setTimeout(() => {
        typingRef.current.innerText = ''
      }, 5000)
    } else {
      chatRef.current.innerHTML += `<div className="chip">
                                      <img src="${gravatarUrl(msg.email)}">
                                        ${msg.username}
                                    </div>
                                    ${msg.message} <br/>`
    }
  }

  useEffect(() => {
    console.log(chatRef);
  }, [chat])

  useEffect(() => {
    if (window.location.host === '127.0.0.1:8000') {
      var newWs = new WebSocket(`ws://${window.location.host}/ws`)
    } else {
      var newWs = new WebSocket(`ws://127.0.0.1:8000/ws`)
    }

    setWs(newWs)


  }, [])

  useEffect(() => {
    if (ws) {
      ws.addEventListener('message', handleBroadcasted)
    }
  }, [ws])


  return (
    <>
    <header>
      <nav>
        <div className="nav-wrapper">
          <a href="/" className="brand-logo right">
            Simple Chat
          </a>
        </div>
      </nav>
    </header>
    <main>
      <div className="row">
        <div className="col s12">
          <div className="card horizonral">
            <div ref={chatRef} id="chat-message" className="card-content">
            </div>
          </div>
        </div>
      </div>
    </main>
    <span><p ref={typingRef} ></p></span>
    {
      joined ?
      <div className="row">
          <div className="input-field col s8">
              <input type="text" value={message} onChange={(e) => {typing();setMessage(e.target.value)}} />
          </div>
          <div className="input-field col s4">
              <button className="waves-effect waves-light btn" onClick={() => send()}>
                  <i className="material-icons right">chat</i>
                  Send
              </button>
          </div>
      </div>
      :
      <div className="row">
        <div className="input-field col s8">
            <input type="email" value={email} placeholder="Email" onChange={(e) => setEmail(e.target.value)}/>
        </div>
        <div className="input-field col s8">
            <input type="text" value={userName} placeholder="Username" onChange={(e) => setUserName(e.target.value)}/>
        </div>
        <div className="input-field col s4">
            <button className="waves-effect waves-light btn" onClick={() => join()}>
                <i className="material-icons right">done</i>
                Join
            </button>
        </div>
    </div>

  }
    </>
  );
}

export default App;
