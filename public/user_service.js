class UserService {
    getPublicContent() {
        return axios.get(API_AUTH_URL + 'all');
    }

    getUserBoard() {
        return axios.get(API_AUTH_URL + 'users' );
    }
    createUser(user){
        return axios.post(API_AUTH_URL + 'user',user)
    }
    disableUser(userId){
        return axios.put(API_AUTH_URL + 'user/disable/'+userId)
    }
    resetPassword(userId){
        return axios.put(API_AUTH_URL + 'user/reset_password'+userId)
    }
    getModeratorBoard() {
        return axios.get(API_AUTH_URL + 'mod' );
    }

    getAdminBoard() {
        return axios.get(API_AUTH_URL + 'admin');
    }
}