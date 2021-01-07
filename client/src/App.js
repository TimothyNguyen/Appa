
import React from 'react';
import jwt_decode from "jwt-decode";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import { Provider } from 'react-redux';

import { Container } from 'react-bootstrap';

import Login from './pages/Login';
import Signup from './pages/Signup';

import NavBar from './components/CustomNavbar';
// import Footer from './components/Footer';
import PrivateRoute from './components/PrivateRoute';
import setAuthToken from './utils/setAuthToken';
import { milisecondsToSeconds } from './utils/dateTime';

import store from "./redux/store";

import './styles/App.css';

// Check for token to keep user logged in
/*
if (localStorage.jwtToken) {
  const token = localStorage.jwtToken;
  setAuthToken(token);
  const decoded = jwtDecode(token);
  store.dispatch(setCurrentUser(decoded));
  const currentTime = milisecondsToSeconds(Date.now());
  if (decoded.exp < currentTime) {
    store.dispatch(logoutUser());
    window.location.href = './login';
  }
}
*/

function App() {
  return (
    <div style={{ minHeight: '100vh', background: '#eeeeee' }}>
      <Provider store={store}> 
        <NavBar />
        <main>
          <Switch>
            <Route path="/login" component={Login} />
            <Route path="/signup" component={Signup} />
            <>
              <Container
                  style={{
                    marginTop: '25px',
                    background: '#ffffff',
                  }}>
            
              </Container>
            </>
          </Switch>
        </main>
      </Provider>
    </div>
  );
}

export default App;
