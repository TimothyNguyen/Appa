import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

const routeUserToComponent = (userType) => {

};

const Home = ({ auth }) => {
    const userType = auth.user.user_type;
    return <>{routeUserToComponent(userType)}</>;
  };
  
  Home.propTypes = {
    // eslint-disable-next-line react/forbid-prop-types
    auth: PropTypes.object.isRequired,
  };
  
  const mapStateToProps = (state) => ({
    auth: state.auth,
  });
  
  export default connect(mapStateToProps)(Home);