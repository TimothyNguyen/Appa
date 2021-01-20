import axios from 'axios';
import jwtDecode from 'jwt-decode';
import { message } from 'antd';
import setAuthToken from '../../utils/setAuthToken';

import { GET_ERRORS, SET_CURRENT_USER, USER_LOADING } from './types';

/**
 * 
 * @param {*} decoded 
 * Helps user login
 */
export const setCurrentUser = (decoded) => ({
    type: SET_CURRENT_USER,
    payload: decoded
});

/**
 * Action to load the use
 */
export const setUserLoading = () => ({
    type: USER_LOADING
});

/**
 * 
 * @param {*} userData 
 * @param {*} history 
 */
export const registerUser = (userData, history) => (dispatch) => {
    axios
    .post('http://localhost:8000/register', userData)
    .then(() => history.push('/login'))
    .catch((err) => {
        console.log("Hi")
        console.log(err)
        dispatch({
            type: GET_ERRORS,
            payload: err,
        });
    });
};

export const loginUser = (userData, history) => (dispatch) => {
    axios
    .post('http://localhost:8000/login', userData)
    .then((res) => {
        if (res.status === 400) {
            // eslint-disable-next-line no-console
            console.log(res.json());
        }
        // console.log(res.data.data);
        const { token } = res.data.data;
        // Set token to localStorage
        localStorage.setItem('jwtToken', token);
        console.log(token);
        // Set token to auth header
        setAuthToken(token);
        // Decode token to get user data
        const decoded = jwtDecode(token);
        // Set current user
        dispatch(setCurrentUser(decoded));
        // Add possible message
        message.success('Login Successful');
        // Go to the home page
        history.push('/');
    })
    .catch((err) => {
        console.log(err);
        // eslint-disable-next-line no-console
        /*
        console.log(err.response.data);
        if (typeof err.response.data.emailnotfound !== 'undefined') {
            message.error("Email doesn't exist");
        } else if (typeof err.response.data.passwordincorrect !== 'undefined') {
            message.error('Password is incorrect');
        }
        */
        dispatch({
            type: GET_ERRORS,
            payload: err.message,
        });
    });
};

// Log user out
export const logoutUser = () => (dispatch) => {
    // Remove token from local storage
    localStorage.removeItem('jwtToken');
    // Remove auth header for future requests
    setAuthToken(false);
    // Set logout messaage
    message.success('Logout Successful');
    // Set current user to empty object {} which will set isAuthenticated to false
    dispatch(setCurrentUser({}));
};
