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

    return (
        <div>
            <p>Profile!</p>
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