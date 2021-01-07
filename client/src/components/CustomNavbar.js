import React from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

import { Container, Nav, Navbar } from 'react-bootstrap';
import { Button } from 'antd';

function CustomNavbar() {
    // const logout = () => logoutUser();
    const route = window.location.pathname.substr(1);
    const loggedOutNav = (route) => (
        <>
        <Nav.Link href="/login" active={route === 'login'}>
            Login
        </Nav.Link>
        <Nav.Link href="/signup" active={route === 'signup'}>
            Signup
        </Nav.Link>
        </>
    );

    
    return (
      <Navbar bg="light" expand="lg">
        <Container>
          <Navbar.Brand href="/">Appa</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav" className="justify-content-end">
            <Nav className="mr-right">
              {loggedOutNav(route)}
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    );
}

CustomNavbar.propTypes = {
    auth: PropTypes.object.isRequired,
    logoutUser: PropTypes.func.isRequired,
};

const mapStateToProps = (state) => ({
    auth: state.auth,
});
  
export default connect(mapStateToProps, { })(CustomNavbar);