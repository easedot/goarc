new Vue({
    el: '#app',
    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false, // True if email and username have been filled in
        online: false
    },
    created: function() {
        let self = this;
        self.connect()
    },
    methods: {
        connect: function() {
            let self = this;
            this.ws = new WebSocket('ws://' + window.location.host + '/ws');
            this.ws.addEventListener('message', function (e) {
                let msg = JSON.parse(e.data);
                self.chatContent += '<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                    + msg.username
                    + '</div>'
                    + emojione.toImage(msg.message) + '<br/>'; // Parse emojis

                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
            });
            this.ws.addEventListener('close', function (e) {
                console.log("socket close")
                self.online = false
            })
            this.ws.addEventListener('open', function (e) {
                console.log("socket open")
                self.online = true
            })
        },
        waitConnect: async  function(timeout =1000){
            const isOpened = () => (this.ws.readyState === WebSocket.OPEN)
            if (this.ws.readyState !== WebSocket.CONNECTING) {
                return isOpened()
            }
            else {
                const intrasleep = 100
                const ttl = timeout / intrasleep // time to loop
                let loop = 0
                while (this.ws.readyState === WebSocket.CONNECTING && loop < ttl) {
                    await new Promise(resolve => setTimeout(resolve, intrasleep))
                    loop++
                }
                return isOpened()
            }
        },
        send: async function () {
            if (!this.online) {
                this.connect()
            }
            const opened =await this.waitConnect()
            if (opened && this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                            email: this.email,
                            username: this.username,
                            message: $('<p>').html(this.newMsg).text() // Strip out html
                        }
                    ));
                this.newMsg = ''; // Reset newMsg
            }
            else {
                console.log("the socket is closed OR couldn't have the socket in time, program crashed");
            }
        },
        join: function () {
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
        },
        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    }
});