
class AuthService {
    login(user) {
        return axios
            .post(API_URL + 'sign_in', {
                name: user.name,
                password: user.password,
                captcha_id: user.captcha_id,
                captcha_code:user.captcha_code
            })
            .then(response => {
                if (response.data.accessToken) {
                    localStorage.setItem('user', JSON.stringify(response.data));
                }

                return response.data;
            });
    }

    logout() {
        localStorage.removeItem('user');
    }

    register(user) {
        return axios.post(API_URL + 'sign_up', {
            name: user.name,
            email: user.email,
            captcha_id: user.captcha_id,
            captcha_code:user.captcha_code,
            password: user.password
        }).then(
            response => {
                if (response.data.accessToken) {
                    localStorage.setItem('user', JSON.stringify(response.data));
                }

                return response.data;
            }
        );
    }
}
