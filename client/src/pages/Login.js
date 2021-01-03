import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Form, Input, Button, message } from 'antd';


const Login = ({ auth, loginUserAction, history, errors }) => {

    return (
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
            ]}
          >
            <Input.Password />
          </Form.Item>
        </Form>
        </div>
    )

};

export default Login;