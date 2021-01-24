const API_URL = 'http://localhost:8080/';
const API_AUTH_URL = 'http://localhost:8080/api/v1/';

const USER_STATE_APPROVE=1
const USER_STATE_DISABLED=2

const UserState={
    1:"正常",
    2:"禁用"
}
const USER_TYPE_VENDOR=1
const USER_TYPE_ADMIN=2
const USER_TYPE_OPER=3

const UserType={
    1:"供应商",
    2:"管理员",
    3:"操作员"
}


class User {
    constructor(name, email, password,captcha_id,captcha_code) {
        this.name = name;
        this.email = email;
        this.password = password;
        this.captcha_id =captcha_id;
        this.captcha_code = captcha_code;
    }
}

Vue.use(httpVueLoader);

axios.interceptors.request.use(function (config) {
    // const token = store.getState().session.token;
    let user = JSON.parse(localStorage.getItem('user'));
    if (user && user.accessToken) {
        config.headers.Authorization = 'Bearer ' + user.accessToken;
    }
    return config;
});
