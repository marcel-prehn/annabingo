import React from "react";
import {Col} from "react-bootstrap";

interface FieldProps {
    text: string,
}

export const Bingofield: React.FC<FieldProps> = props => {
    const {text} = props
    const [active, setActive] = React.useState(false)

    return (
        <Col className={`bingofield ${active ? "activefield" : ""}`}
             onClick={() => setActive(!active)}>{text}
        </Col>
    )
}
