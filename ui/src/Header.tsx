import React from "react";
import {Nav, Navbar} from "react-bootstrap";

export const Header = () => {
  return (
    <Navbar className="navbar-custom">
      <Navbar.Brand href="#">ANNABINGO!</Navbar.Brand>
      <Nav className="mr-auto">
        <Nav.Link href="#create">+ Neu</Nav.Link>
        <Nav.Link href="#search">o Suchen</Nav.Link>
        <Nav.Link className={"justify-content-end"} href="#impressum">Impressum</Nav.Link>
      </Nav>
      <div>
        <a href="https://www.buymeacoffee.com/MarcelPrehn" target="_blank">
          <img className={"coffee"} src="https://cdn.buymeacoffee.com/buttons/default-black.png" alt="Buy Me A Coffee"/>
        </a>
      </div>
    </Navbar>
  )
}