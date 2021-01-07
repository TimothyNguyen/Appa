import axios from 'axios';
import setAuthToken from '../../utils/setAuthToken';

import { GET_ERRORS, UPDATE_CURRENT_USER } from './types';

// Set logged in user
export const updateCurrentUser = (payload) => ({
    type: UPDATE_CURRENT_USER,
    payload,
});

export const updateProfileInfo = (data) => (dispatch) => {
    axios
        .patch('/profile/update_details', data)
        .then((res) => {
            const { newData, user } = data;
            // update currentuser
            dispatch(updateCurrentUser(Object.assign(user, newData)));
            const { token } = res.data;
            // Set token to localStorage
            localStorage.setItem('jwtToken', token);
            // Set token to Auth header
            setAuthToken(token);
        })
        .catch((err) =>
            dispatch({
                type: GET_ERRORS,
                payload: err.response.data,
            })
        );
};

export const updateProfilePic = () => {};
  
