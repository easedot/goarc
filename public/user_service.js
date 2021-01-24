class UserService {
    getPublicContent() {
        return axios.get(API_AUTH_URL + 'all');
    }
    getUserBoard(param) {
        return axios.get(API_AUTH_URL + 'users',{params:param} );
    }
    createUser(user){
        return axios.post(API_AUTH_URL + 'user',user)
    }
    enableUser(userId){
        return axios.put(API_AUTH_URL + 'user/enable/'+userId)
    }
    disableUser(userId){
        return axios.put(API_AUTH_URL + 'user/disable/'+userId)
    }
    resetPassword(userId){
        return axios.put(API_AUTH_URL + 'user/reset_password/'+userId)
    }
    changePassword(user){
        return axios.put(API_AUTH_URL + 'user/change_password',user)
    }
    getModeratorBoard() {
        return axios.get(API_AUTH_URL + 'mod' );
    }
    getAdminBoard() {
        return axios.get(API_AUTH_URL + 'admin');
    }
}