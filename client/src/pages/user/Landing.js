import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Button } from 'antd';

const Landing = ({ auth }) => {
  const { name } = auth.user;
  return (
    <div
        className="site-layout-content"
        style={{
            marginLeft: '5px',
            marginRight: '5px',
        }}>
        <h2>Hello {name}!</h2>
        <p>
        Thanks for being part of Appa. 
        We are still developing the application
        and will have further updates soon.
        </p>
        <Button
            type="primary"
            size="large"
            href="/profile"
            style={{ margin: '8px' }}>
            Profile
        </Button>
     </div> 
  );
}

Landing.propTypes = {
    // eslint-disable-next-line react/forbid-prop-types
    auth: PropTypes.object.isRequired,
  };
  
  const mapStateToProps = (state) => ({
    auth: state.auth,
  });
  
  export default connect(mapStateToProps)(Landing);