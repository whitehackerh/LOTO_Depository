<template>
    <div>
        <HeaderComponent />
        <div class="registerForm">
            <label>User Name<input type="text" v-model="user_name"></label><br>
            <label>Password<input type="text" v-model="password"></label><br>
            <label>Mail Address<input type="text" v-model="mail_address"></label><br>
            <button @click="createUser()">Register</button>
        </div>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';

export default {
    name: 'UserRegistrationPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            user_name: '',
            password: '',
            mail_address: '',
        }
    },
    methods: {
        async createUser() {
            await axios.post("/createUser", {body: {user_name: this.user_name,
            password: this.password, mail_address: this.mail_address}})
        }
    }
}
</script>

<style>
.registerForm {
    margin: 0 auto;
    margin-top: 50px;
}
label {
    color: white;
}
button {
    text-align: center;
}
</style>
