import React from "react";
import {Button, Col, Container, Row} from "react-bootstrap";
import {CreateField} from "./CreateField";
import {Link} from "react-router-dom";
import {Bingo} from "./Bingo";

export const CreateMatrix = () => {
  const [value, setValue] = React.useState<string[][]>([[], [], [], []])
  const [title, setTitle] = React.useState("")
  const [url, setUrl] = React.useState("")
  const [uuid, setUuid] = React.useState("")

  const setField = (text: string, row: number, cell: number) => {
    value[row][cell] = text
    setValue(value)
  }

  const postMatrix = async () => {
    let bingo: Bingo = {
      title: title,
      fields: value
    }
    const host = "http://localhost:3000/#/view/"
    let response = await fetch('/api/create', {
      method: 'post',
      body: JSON.stringify(bingo),
      headers: {'Content-Type': 'application/json'}
    })
    setUuid(await response.text())
    setUrl(host + uuid)
  }

  return (
    <Container>
      <Row>
        <Col><input value={title} onChange={e => setTitle(e.target.value)} placeholder={"Titel"}
                    className={"form-control bingotitle mt-3"}/></Col>
      </Row>
      <Row>
        <Col><CreateField text={value[0][0]} onChange={event => (setField(event.target.value, 0, 0))}/></Col>
        <Col><CreateField text={value[0][1]} onChange={event => (setField(event.target.value, 0, 1))}/></Col>
        <Col><CreateField text={value[0][2]} onChange={event => (setField(event.target.value, 0, 2))}/></Col>
        <Col><CreateField text={value[0][3]} onChange={event => (setField(event.target.value, 0, 3))}/></Col>
      </Row>
      <Row>
        <Col><CreateField text={value[1][0]} onChange={event => (setField(event.target.value, 1, 0))}/></Col>
        <Col><CreateField text={value[1][1]} onChange={event => (setField(event.target.value, 1, 1))}/></Col>
        <Col><CreateField text={value[1][2]} onChange={event => (setField(event.target.value, 1, 2))}/></Col>
        <Col><CreateField text={value[1][3]} onChange={event => (setField(event.target.value, 1, 3))}/></Col>
      </Row>
      <Row>
        <Col><CreateField text={value[2][0]} onChange={event => (setField(event.target.value, 2, 0))}/></Col>
        <Col><CreateField text={value[2][1]} onChange={event => (setField(event.target.value, 2, 1))}/></Col>
        <Col><CreateField text={value[2][2]} onChange={event => (setField(event.target.value, 2, 2))}/></Col>
        <Col><CreateField text={value[2][3]} onChange={event => (setField(event.target.value, 2, 3))}/></Col>
      </Row>
      <Row>
        <Col><CreateField text={value[3][0]} onChange={event => (setField(event.target.value, 3, 0))}/></Col>
        <Col><CreateField text={value[3][1]} onChange={event => (setField(event.target.value, 3, 1))}/></Col>
        <Col><CreateField text={value[3][2]} onChange={event => (setField(event.target.value, 3, 2))}/></Col>
        <Col><CreateField text={value[3][3]} onChange={event => (setField(event.target.value, 3, 3))}/></Col>
      </Row>
      <Row>
        <div className="ml-auto mr-auto mb-3">
          <Button onClick={postMatrix}>Create Annabingo!</Button>
          <Link className={`bingolink ${url.length === 0 ? "displaynone" : ""}`} to={`/view/${uuid}`}>Open
            Annabingo</Link>
        </div>
      </Row>
    </Container>
  )
}