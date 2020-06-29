import React from 'react';
import {
  Navbar,
  NavbarBrand,
} from 'reactstrap';

const NavBar = () => {

  return (
    <div>
      <Navbar color="dark" dark expand="md">
        <NavbarBrand href="/">Chat App</NavbarBrand>
      </Navbar>
    </div>
  );
}

export default NavBar;
