import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import axios from 'axios';
import '../styles/App.css';

import { Row, Col, Image, Form } from 'react-bootstrap';
import {
    Button,
    Descriptions,
    Divider,
    Modal,
    Space,
    Typography,
    Upload,
} from 'antd';
import ImgCrop from 'antd-img-crop';
import { UploadOutlined } from '@ant-design/icons';

const { Paragraph } = Typography;

const center = {
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
};

const Profile = ({ auth, history }) => {

  const [imageURI, setImageURI] = useState('favicon.ico');
  const [showEditModal, setShowEditModal] = useState(true);
  
  const { id } = auth.user;

  console.log(auth.user.user);

  return (
    <div className="site-layout-content">
      <h4>Profile</h4>
      <Row className="justify-content-md-center">
        <Col md={4} style={{ ...center, marginTop: '12px' }}>
          <div>
            <Row style={{ ...center, marginTop: '12px' }}>
              <Image
                src={imageURI}
                alt="test-img"
                style={{ minWidth: '200px' }}
                roundedCircle
              />
            </Row>
          </div>
        </Col>
      </Row>
      <Descriptions title="User Info">
        <Descriptions.Item label="Name">{auth.user.user.name}</Descriptions.Item>
        <Descriptions.Item label="Email">{auth.user.user.email}</Descriptions.Item>
        <Descriptions.Item label="Phone Number">
          {auth.user.user.phone_number}
        </Descriptions.Item>
        <Descriptions.Item label="GitHub Username">
          {auth.user.user.github_username}
        </Descriptions.Item>
        <Descriptions.Item label="LinkedIn">
          {auth.user.user.linkedin}
        </Descriptions.Item>
      </Descriptions>
      <Divider />
      <Descriptions title="Actions" />
      <Space size="middle">
        <Button type="primary">
          Edit Profile
        </Button>
        <Button type="primary" danger>
          Delete Account
        </Button>
      </Space>
    </div>
  )
}

Profile.propTypes = {
  // eslint-disable-next-line react/forbid-prop-types
  auth: PropTypes.object.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired,
  }).isRequired,
};

const mapStateToProps = (state) => ({
  auth: state.auth,
});

export default connect(mapStateToProps)(Profile);