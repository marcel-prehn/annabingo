import React from "react";
import {Container} from "react-bootstrap";
import {Link} from "react-router-dom";
import bingo from "./bingo-image-unsplash.jpg"

export const Home = () => {
    return (
        <Container className="text-center">
            <h1 className="mt-5">Annabingo Annabongo</h1>
            <div>
                Klicke auf <Link to="/create">Create</Link> um ein neues Annabingo anzulegen mit deinen eigenen Aussagen.<br/>
                Nach dem erfolgreichen Anlegen, wird ein Link generiert den du mit deinen Freunden teilen kannst.<br/>
                Enjoy Annabingo!
                <div><img className="img-fluid mt-5" src={bingo} alt="Annabingo Splash"/></div>
            </div>
        </Container>
    )
}

