import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Landing from './user/Landing';

// Utilize this userType later on if we want to have an admin vs members page
const routeUserToComponent = (userType) => {
  // console.log(userType);
  return <Landing />
};

const Home = ({ auth }) => {
  console.log(auth);
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