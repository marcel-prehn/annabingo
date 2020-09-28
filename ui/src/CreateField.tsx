import React, {FC} from "react";
import {Col} from "react-bootstrap";

interface CreateFieldProps {
    text: string,

    onChange(event: any): void,
}

export const CreateField: FC<CreateFieldProps> = (props) => {
    const {text, onChange} = props

    return (
        <Col><textarea className="createfield" onChange={onChange} value={text}/></Col>
    )
}
