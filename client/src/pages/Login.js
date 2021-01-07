import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Form, Input, Button, message } from 'antd';

import Background from '../assets/leaves.jpg';

const Login = ({ auth, loginUserAction, history, errors }) => {

    return (
        <div
            style={{
                minHeight: '100vh',
                backgroundImage: `url(${Background})`,
                backgroundSize: 'cover',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
            }}>
        <div
            style={{
            width: '550px',
            padding: '25px',
            background: 'rgba(255, 255, 255, 1.0)',
            }}>
            <h4>Log in</h4>
            <Form
            layout="vertical"
            name="basic"
            initialValues={{
                remember: true,
            }}>
                <Form.Item
                label="Email"
                name="email"
                rules={[
                {
                    required: true,
                    message: 'Please input your email!',
                },
                ]}
            >
                <Input />
            </Form.Item>
            <Form.Item
                label="Password"
                name="password"
                rules={[
                {
                    required: true,
                    message: 'Please input your password!',
                },
                ]}>
                <Input.Password />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                    Submit
                    </Button>
                </Form.Item>
                </Form>
            </div>
        </div>
    )

};


Login.propTypes = {
    loginUserAction: PropTypes.func.isRequired,
    history: PropTypes.shape({
      push: PropTypes.func.isRequired,
    }).isRequired,
    // eslint-disable-next-line react/forbid-prop-types
    auth: PropTypes.object.isRequired,
    // eslint-disable-next-line react/forbid-prop-types
    errors: PropTypes.object.isRequired,
  };
  
  const mapStateToProps = (state) => ({
    auth: state.auth,
    errors: state.errors,
  });
  
  export default connect(mapStateToProps, { })(Login);