import axios from 'axios';
import jwtDecode from 'jwt-decode';

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
        .post('/auth/register', userData)
        .then(() => history.push('/login'))
        .catch((err) => {
            dispatch({
                type: GET_ERRORS,
                payload: err.response.data,
            });
        });
};
