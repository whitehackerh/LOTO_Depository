<template>
    <HeaderComponent />
    <div id="app">
        <input type="text" v-model="user_name" placeholder="Username">
        <input type="password" v-model="password" placeholder="Password">
    </div>
    <div>
        <button @click="Login()">Login</button>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'LoginPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            user_name: '',
            password: '',
            results: '',
            user: {}
        }
    },
    methods: {
        async Login() {
            let respon = await axios.post("/login", {body: {user_name: this.user_name, password: this.password}});
            if (respon.data.result == "success") {
                this.$store.dispatch("auth", {
                        user_id: respon.data.user_id,
                        user_name: respon.data.token,
                        mail_address: respon.data.mail_address,
                        admin_flag: respon.data.admin_flag,
                        token: respon.data.token
                })
                this.$router.push('/');
            } else {
                this.results = "Failed"
            }
        }
    }
};
</script>