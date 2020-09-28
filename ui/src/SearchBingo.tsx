import React, {useEffect, useState} from "react";
import {Col, Container, ListGroup, Row} from "react-bootstrap";
import {Bingo} from "./Bingo";

export const SearchBingo = () => {
  const [query, setQuery] = useState("")
  const [result, setResult] = useState<Array<Bingo>>()

  useEffect(() => {
    query.length === 0? setResult([]) :
    fetch(`/api/search/${query}`)
      .then(response => response.json())
      .then(data => setResult(data))
  }, [query])

  const open = (uuid?: string) => {
    window.location.href = `#/view/${uuid}`
  }

  return (
    <Container>
      <Row>
        <Col>
          <input onChange={event => setQuery(event.target.value)} className={"form-control mt-3"}
                 placeholder={"Annabingo Titel ..."}/>
        </Col>
      </Row>
      <Row>
        <ListGroup className="m-3 bingosearch">
          {
            result?.map((bingo) => (
              <ListGroup.Item key={bingo.title} action onClick={() => open(bingo.uuid)}>{bingo.title}</ListGroup.Item>
            ))
          }
        </ListGroup>
      </Row>
    </Container>
  )
}