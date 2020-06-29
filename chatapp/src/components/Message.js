import React from 'react';
import {Row, Col, Card} from 'reactstrap'


const Message = ({msg}) => {

  return (
    <Row>
      <Col md="12">
        <Card body inverse style={{ marginTop: "4px",backgroundColor: '#333', borderColor: '#333', height: "100px" }}>
          <div className="d-flex">
            <p className="px-1">{msg.sender.userName}</p>
            <p className="px-1">{msg.timestamp}</p>
          </div>
          <p className="px-1">{msg.message}</p>
        </Card>
      </Col>
    </Row>
  )
}

export default Message
