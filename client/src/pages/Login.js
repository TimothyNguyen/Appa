import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Form, Input, Button, message } from 'antd';

import { loginUser } from '../redux/actions/authActions';

import Background from '../assets/leaves.jpg';

const Login = ({ auth, loginUserAction, history, errors }) => {
    
    const onFinish = (values) => {
        loginUserAction(values, history);
    };
    
    const onFinishFailed = (errorInfo) => {
        // eslint-disable-next-line no-console
        message.error('Cannot login. Server may not be running');
        console.log('Failed:', errorInfo, errors);
    };
    
    if (auth.isAuthenticated) {
        history.push('/');
    }

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
                }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}>
                <Form.Item
                    label="Email"
                    name="Email"
                    rules={[
                    {
                        required: true,
                        message: 'Please input your email!',
                    },
                    ]}>
                    <Input />
                </Form.Item>
                <Form.Item
                    label="Password"
                    name="Password"
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
  
  export default connect(mapStateToProps, { loginUserAction: loginUser })(Login);