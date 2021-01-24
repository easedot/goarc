
    /*1:供应商 2:管理员 3:运营人员*/
const authService=new AuthService()
const userService=new UserService()

const user = JSON.parse(localStorage.getItem('user'));
const initialState = user
    ? { status: { loggedIn: true }, user }
    : { status: { loggedIn: false }, user: null };

const auth = {
    namespaced: true,
    state: initialState,
    mutations: {
        loginSuccess(state, user) {
            state.status.loggedIn = true;
            state.user = user;
        },
        loginFailure(state) {
            state.status.loggedIn = false;
            state.user = null;
        },
        logout(state) {
            state.status.loggedIn = false;
            state.user = null;
        },
        registerSuccess(state) {
            state.status.loggedIn = true;
            state.user = user;
        },
        registerFailure(state) {
            state.status.loggedIn = false;
            state.user = null;
        }
    },
    actions: {
        login:function (context,user){
            console.info("auth login call",user)
            return authService.login(user).then(
                user => {
                    context.commit('loginSuccess', user);
                    return Promise.resolve(user);
                },
                error => {
                    context.commit('loginFailure');
                    return Promise.reject(error);
                }
            );
        },
        logout:function (context){
            authService.logout();
            context.commit('logout')
        },
        register:function (context,user){
            return authService.register(user).then(
                response => {
                    context.commit('registerSuccess');
                    return Promise.resolve(response.data);
                },
                error => {
                    context.commit('registerFailure');
                    return Promise.reject(error);
                }
            );
        }
    },
    // sign_up:function (name,password,captcha_id,captcha_code){
    //     return axios
    //         .post( '/sign_up', {
    //             name: name,
    //             password: password,
    //             captcha_id,
    //             captcha_code
    //         })
    //         .then(response => {
    //             console.info("sign_up",response)
    //             if (response.data.accessToken) {
    //                 localStorage.setItem('user', JSON.stringify(response.data));
    //             }
    //             return response.data;
    //         })
    //
    // },
    // sign_in:function (name,password,captcha_id,captcha_code){
    //     return axios
    //     .post( '/sign_in', {
    //         name: name,
    //         password: password,
    //         captcha_id,
    //         captcha_code
    //     })
    //     .then(response => {
    //         console.info("sign_in",response)
    //         if (response.data.accessToken) {
    //             localStorage.setItem('user', JSON.stringify(response.data));
    //         }
    //         return response.data;
    //     }).catch(function (error) {
    //         console.log(error.response.status) // 401
    //         console.log(error.response.data.error) //Please Authenticate or whatever returned from server
    //         if(error.response.status===401){
    //             //redirect to login
    //         }
    //     })
    // }
}
const store = new Vuex.Store({
    state: {
        count: 0
    },
    modules: {
        auth
    },
    mutations: {
        increment (state) {
            state.count++
        }
    }
})
const upload_file = {
    props:["name","value"],
    data:function (){
        return {
            files:[],
            upload_files:[],
            uploadPercentage:0
        }
    },
    created:function (){
        if (this.value){
            console.info("init upload value",this.value)
            this.upload_files=this.value.split(",")
        }
    },
    template:"#upload_file",
    computed:{
        file_names:function (){
            if (this.value){
                this.upload_files = this.value.split(",")
            }else{
                this.upload_files=[]
            }
            return this.upload_files
        }
    },
    methods:{
        getFile(file){
            let url= this.uploadUrl(file)
            axios.get(url,{responseType: 'blob'}).then(
                response=>{
                    console.info("get file ",response)
                    const url = window.URL.createObjectURL(new Blob([response.data]));
                    const link = document.createElement('a');
                    link.href = url;
                    link.setAttribute('download', file); //or any other extension
                    document.body.appendChild(link);
                    link.click();
                }
            )
        },
        uploadUrl(file){
           return "/api/v1/upload_file/"+file
        },
        addFiles(){
            this.$refs.files.click();
        },
        removeFile( key ){
            let self = this
            console.info("delete")
            let url='/api/v1/upload_file/'+this.upload_files[key]
            console.info("delete url",url)
            axios.delete(url)
                .then(function(resp){
                    console.log('Delete upload file SUCCESS!!',resp);
                    self.upload_files.splice( key, 1 );
                    let vl=self.upload_files.join(",")
                    self.$emit('input', vl)
                }).catch(function(resp){
                console.log('Delete upload file FAILURE!!',resp);
            });
        },
        handleFilesUpload(){
            let uploadedFiles = this.$refs.files.files;
            for( var i = 0; i < uploadedFiles.length; i++ ){
                this.files.push( uploadedFiles[i] );
            }
            this.submitFiles()
        },
        submitFiles(){
            let self = this
            let formData = new FormData();
            formData.append("name",this.name)
            for( var i = 0; i < this.files.length; i++ ){
                let file = this.files[i];
                formData.append('files[]', file);
            }
            axios.post( '/api/v1/upload_file',
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    },
                    onUploadProgress: function( progressEvent ) {
                        this.uploadPercentage = parseInt(Math.round((progressEvent.loaded/progressEvent.total)*100))
                    }.bind(this)
                }
            ).then(function(resp){
                console.log('SUCCESS!!',resp);
                if (resp.data){
                    console.log("append upload file",self.upload_files,resp.data)
                    self.upload_files=self.upload_files.concat(resp.data)
                    console.log("uploaded file",self.upload_files)
                    let lv=self.upload_files.join(",")
                    self.$emit('input', lv)
                    self.files=[]
                }
            })
                .catch(function(resp){
                    console.log('FAILURE!!',resp);
                    self.files=[]
                });
        }
    }
}

const captcha ={
    props:{
        "captcha_id":String,
        "captcha_code":String,
    },
    data:function (){
        return{
            // captcha_id:"",
            // captcha_code:"",
            captcha_url:""
        }
    },
    created:function (){
        let self = this
        axios.get('/captcha/id')
            .then(function(resp){
                console.log('Captcha id SUCCESS!!',resp.data);
                self.captcha_id = resp.data.id
                console.log("captcha_id",resp.data.id)
                self.captcha_url = '/captcha/img/'+self.captcha_id
                console.log("captcha_url",self.captcha_url)
                self.$emit('update:captcha_id', self.captcha_id)
            }).catch(function(resp){
                console.log('Captcha id FAILURE!!',resp);
            });
    },
    template:"#captcha",
    methods:{
        change_captcha_code:function (){
            console.info("change captcha code",this.captcha_code)
            this.$emit('update:captcha_code',this.captcha_code)
        },
        refresh:function (){

            console.info("refresh img")
            let timestamp = new Date().getTime();
            let queryString = "?t=" + timestamp;
            let url='/captcha/img/'+this.captcha_id+queryString
            console.info("reload imag url:",url)
            this.captcha_url=url
            // let img=document.getElementById('captcha-img')
            // img.src = url
        }
    }
}


const SignIn = {
    data:function(){
        return{
            user: new User('', '', '','',''),
            loading: false,
            message: ''
        }
    },
    template: "#sign_in",
    methods:{
        sign_in:function (){
            if (this.user.name && this.user.password && this.user.captcha_code) {
                this.$store.dispatch('auth/login', this.user).then(
                    () => {
                        this.$router.push('/profile');
                    },
                    error => {
                        this.loading = false;
                        this.message =
                            (error.response && error.response.data) ||
                            error.message ||
                            error.toString();
                    }
                );
            }
        }
    },
    created() {
        if (this.loggedIn) {
            this.$router.push('/profile');
        }
    },
    computed: {
        loggedIn() {
            return this.$store.state.auth.status.loggedIn;
        }
    },
    components: {
        captcha
    },
};

const SignUp = {
    data:function(){
        return{
            user: new User('', '', '','',''),
            submitted: false,
            successful: false,
            message: ''
        }
    },
    template: "#sign_up",
    computed: {
        loggedIn() {
            return this.$store.state.auth.status.loggedIn;
        }
    },
    mounted() {
        if (this.loggedIn) {
            this.$router.push('/profile');
        }
    },
    methods:{
        sign_up:function (){
            console.info("sign_up...",this.user.name)
            this.message = '';
            this.submitted = true;
            // this.$validator.validate().then(isValid => {
            //     if (isValid) {
                    this.$store.dispatch('auth/register', this.user).then(
                        data => {
                            this.message = "注册成功";
                            this.successful = true;
                            this.$router.push({ name: 'vendor', params: { edit_mode: 'create',user:this.user }});
                        },
                        error => {
                            this.message =
                                (error.response && error.response.data) ||
                                error.message ||
                                error.toString();
                            this.successful = false;
                        }
                    );
                // }
            // })
            // auth.sign_up(this.username,this.password,this.captcha_id,this.captcha_code)
            // this.$router.push({ name: 'vendor', params: { edit_mode: 'create' }});
        }
    },
    components: {
        captcha
    },
};

const reset_password = {
    template: "#reset_password"
};

const VENDOR_STATE_INIT=1
const VENDOR_STATE_SUBMITED=2
const VENDOR_STATE_APPROVED=3
const VENDOR_STATE_REJECT=4
const VENDOR_STATE_REWRITE=5

///*1:待审核 2:审核中 3:通过 4:拒绝 5:补充*/
const Profile ={
    data() {
        return {
            message: '',
            oldPassword:"",
            newPassword:""
        };
    },
    computed: {
        currentUser() {
            return this.$store.state.auth.user;
        },
        userType(){
            return UserType[this.currentUser.type]
        },
        userState(){
            return UserState[this.currentUser.state]
        }
    },
    mounted() {
        if (!this.currentUser) {
            this.$router.push('/login');
        }
    },
    methods:{
        savePassword(){

        }
    },
    template: "#profile"

}
const Vendor = {
    props: ['edit_mode','user'],
    data:function(){
        return{
            vendor:{},
            registered_type_options:{
                1:"私企",
                2:"国企",
                3:"事业",
                4:"外资",
                5:"其他"
            },
            vendor_state_options:{
                1:"待审核",
                2:"审核中",
                3:"通过",
                4:"拒绝",
                5:"补充信息"
            },
            vendor_type_options: {
                1: "普通",
                2: "临时"
            },
            bool_options: {
                1: "是",
                2: "否"
            },
            sucesses_files:"",
            qualification_files:[],
            message:""
        }
    },
    created:function (){
            this.getVendor()
    },
    mounted() {
        if (!this.vendor){
            this.message="blank"
        }
    },
    template: "#vendor",
    computed:{
        vendorState(){
            return this.vendor_state_options[this.vendor.state]
        },
        isShow:function (){
            return this.edit_mode==="show" || this.edit_mode===""
        },
        isCreate:function (){
            return this.edit_mode==="create"
        },
        isUpdate:function (){
            return this.edit_mode==="update"
        },
        allowCreate(){
            return this.isShow && !this.vendor.name
        },
        allowEdit(){
            return this.vendor.state === VENDOR_STATE_INIT ||
                   this.vendor.state === VENDOR_STATE_REWRITE
        },
        currentUser:function() {
            return this.$store.state.auth.user;
        },
        showVendorBoard() {
            if (this.currentUser && this.currentUser.roles) {
                return this.currentUser.roles.includes('ROLE_VENDOR');
            }

            return false;
        },

    },
    methods:{
        saveVendor:function (){
            console.log("do task before submit");
            if (this.$refs.form.checkValidity()) {
                if (this.edit_mode==="update") {
                    this.updateVendor()
                }
                if (this.edit_mode==="create") {
                    this.createVendor()
                }
            } else {
                this.$refs.form.reportValidity();
            }
        },
        toCreate:function (){
            this.$router.push({ name: 'vendor', params: { edit_mode: 'create' }});
        },
        editVendor:function (){
            this.$router.push({ name: 'vendor', params: { edit_mode: 'update' }});
        },
        unSaveVendor(){
            this.$router.push({ name: 'vendor', params: { edit_mode: 'show' }});
        },
        submitVendor:function (){
            axios.post( '/api/v1/vendor', this.vendor)
                .then(response => {
                    console.info("create vendor", response)
                    this.$router.push({ name: 'vendor', params: { edit_mode: 'show' }});
                })
        },
        createVendor:function (){
            axios.post( '/api/v1/vendor', this.vendor)
            .then(response => {
                console.info("create vendor", response)
                this.$router.push({ name: 'vendor', params: { edit_mode: 'show' }});
            })
        },
        updateVendor:function (){
            axios.put( '/api/v1/vendor', this.vendor)
            .then(response => {
                console.info("update vendor", response)
                this.$router.push({ name: 'vendor', params: { edit_mode: 'show' }});
            })
        },
        getVendor:function (){
            axios.get( '/api/v1/vendors', {
            })
            .then(response => {
                console.info("vendor",response)
                if (response.data && response.data.length>0){
                    this.vendor=response.data[0]
                }else{
                    //当没有供应上记录时进入创建模式
                    // this.$router.push({ name: 'vendor', params: { edit_mode: 'create' }});
                }
                console.info("set this vendor:",this.vendor)
            }).catch(function (error) {
                console.log(error.response) // 401
                console.log(error.response.data.error) //Please Authenticate or whatever returned from server
            })
        }
    },
    components: {
        'upload-file': upload_file,
    }
};

const Home={
    data() {
        return {
            content: "火币供应商管理系统"
        };
    },
    mounted() {

    },
    template: "#home"
}
const BoardAdmin={
    data() {
        return {
            content: {'message':'Home'},
            user: new User("","","","","")
        };
    },
    computed:{
    },
    mounted() {
        this.loadData()
    },
    methods:{
        loadData(){
            userService.getUserBoard().then(
                response => {
                    this.content = response.data;
                },
                error => {
                    this.content =
                        (error.response && error.response.data) ||
                        error.message ||
                        error.toString();
                }
            );
        },
        userState(state){
            return UserState[state]
        },
        createUser(){
            userService.createUser(this.user).then(
                response => {
                    this.loadData();
                },
            )
        },
        disableUser(userId){
            userService.disableUser(userId).then(
                response => {
                    this.loadData();
                },
            )
        },
        resetPassword(userId){
            userService.resetPassword(userId).then(
                response => {
                    //ths.loadData();
                },
            )
        }
    },
    template: "#board_admin"

}

const BoardUser={
    data() {
        return {
            content: {'message':'Home'},
            page: 0,
            user: new User("","","","","")
        };
    },
    mounted() {
        this.loadData()
    },
    template: "#board_user",
    methods:{
        prevPage(){
            this.page=this.page-1
            this.loadData()
        },
        nextPage(){
            this.page=this.page+1
            this.loadData()
        },
        showDetail(userId){
            this.$router.push({ name: 'vendor', params: { edit_mode: 'show' }});
        },
        loadData(){
            userService.getUserBoard().then(
                response => {
                    this.content = response.data;
                },
                error => {
                    this.content =
                        (error.response && error.response.data) ||
                        error.message ||
                        error.toString();
                }
            );

        },
        userState(state){
            return UserState[state]
        },
        disableUser(userId){
            userService.disableUser(userId).then(
                response => {
                    this.loadData();
                },
            )
        },
        resetPassword(userId){
            userService.resetPassword(userId).then(
                response => {
                    //ths.loadData();
                },
            )
        }
    }
}

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/home',
        component: Home
    },
    {
        path: '/login',
        component: SignIn
    },
    {
        path: '/profile',
        name: 'profile',
        component: Profile
    },
    {
        path: '/admin',
        name: 'admin',
        component: BoardAdmin
    },
    {
        path: '/oper',
        name: 'oper',
        component: BoardUser
    },
    {
        path: '/sign_up',
        component: SignUp
    },
    {   path: '/vendor/:edit_mode',
        name:"vendor",
        component: Vendor,
        props: true
    },
    {   path: '/reset_password',
        component: reset_password
    }
]

const router = new VueRouter({
    // mode: 'history',
    routes // short for `routes: routes`
})

router.beforeEach((to, from, next) => {
    const publicPages = ['/login', '/sign_up', '/home','/'];
    const authRequired = !publicPages.includes(to.path);
    const loggedIn = localStorage.getItem('user');

    // trying to access a restricted page + not logged in
    // redirect to login page
    console.info("authRequired:",authRequired,"loggedIn:",loggedIn)
    if (authRequired && !loggedIn) {
        next('/login');
    } else {
        next();
    }
});


const app = new Vue({
    router,
    store,
    el: '#app',
    template: "#layout",
    computed:{
        currentUser() {
            return this.$store.state.auth.user;
        },
        showAdminBoard() {
            if (this.currentUser && this.currentUser.roles) {
                return this.currentUser.roles.includes('ROLE_ADMIN');
            }

            return false;
        },
        showOperBoard() {
            if (this.currentUser && this.currentUser.roles) {
                return this.currentUser.roles.includes('ROLE_OPER');
            }

            return false;
        },
        showVendorBoard() {
            if (this.currentUser && this.currentUser.roles) {
                return this.currentUser.roles.includes('ROLE_VENDOR');
            }

            return false;
        }
    },
    methods:{
        logOut:function (){
          this.$store.dispatch('auth/logout');
          this.$router.push('/login');
      }
    }
})
