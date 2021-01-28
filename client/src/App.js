
import React from 'react';
import jwtDecode from "jwt-decode";
import { BrowserRouter as Router, Route, Switch, useLocation } from "react-router-dom";
import { Provider } from 'react-redux';
import { Container } from 'react-bootstrap';

import Home from './pages/Home';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Profile from './pages/Profile';

import NavBar from './components/CustomNavbar';
import Footer from './components/Footer';
import PrivateRoute from './components/PrivateRoute';
import setAuthToken from './utils/setAuthToken';
import { milisecondsToSeconds } from './utils/dateTime';
import ScrollReveal from './utils/ScrollReveal';
import LayoutDefault from './layouts/LayoutDefault';


import store from "./redux/store";
import { setCurrentUser, logoutUser } from './redux/actions/authActions';

import './styles/App.css';

// Check for token to keep user logged in

if (localStorage.jwtToken && localStorage.jwtToken !== "undefined") {
  var retrievedObj = localStorage.getItem('jwtToken');
  var res = JSON.parse(retrievedObj);
  // console.log(res)
  const { access_token } = res.data
  setAuthToken(access_token);
  const decoded = jwtDecode(access_token);
  // console.log(decoded);
  store.dispatch(setCurrentUser(decoded));
  const currentTime = milisecondsToSeconds(Date.now());
  if (decoded.exp < currentTime) {
    store.dispatch(logoutUser());
    window.location.href = './login';
  }

}

function App() {
  const childRef = React.useRef();
  let location = useLocation();

  React.useEffect(() => {
    document.body.classList.add('is-loaded')
    // childRef.current.init();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [location]);

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
                  <PrivateRoute path="/" component={Home} exact/>
                  <PrivateRoute path="/profile" component={Profile} />
                </Container>
                <Footer />
              </>
            </Switch>
          </main>
        </Provider>
      </div>
  );
}

export default App;
