import React, {useState} from "react";
import {Col, Container, Row} from "react-bootstrap";
import {Bingofield} from "./Bingofield";
import {useParams} from "react-router-dom";
import {Bingo} from "./Bingo";

export const Matrix = () => {
  const {uuid} = useParams() as any
  const [bingo, setBingo] = useState<Bingo>()

  React.useEffect(() => {
    fetch(`/api/view/${uuid}`)
      .then(response => response.json())
      .then(data => setBingo(data))
  }, [uuid])

  return (
    <Container>
      <Row className="text-center"><Col className="mt-3"><h2>{bingo?.title}</h2></Col></Row>
      <div className="ml-auto mr-auto">
        {bingo?.fields.map((row, i) => {
          return <Row key={i}>
            {row.map((item, j) => {
              return <Bingofield key={`id-${i}-${j}`} text={item}/>
            })}
          </Row>
        })}
      </div>
      <div className="text-center mt-3 mb-3">
        Aussagen werden nacheinander ausgerufen. Wenn du eine Übereinstimmung mit dir selbst erkennst, markierst du das
        Feld.
        Hast du 3 Treffer quer, hochkant oder diagonal rufst du "Annabingo". <br/>
        Hast du ein Annabingo, gewinnst du das Spiel.
      </div>
    </Container>
  )
}