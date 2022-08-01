import Vuex from 'vuex'

const store = new Vuex.Store({
    state() {
        return {
            user_id: '',
            user_name:'',
            mail_address: '',
            admin_flag: '', 
            token: '',
        }
    },
    mutations: {
        updateUser(state, user) {
            state.user_id = user.user_id;
            state.token = user.token;
            state.user_name = user.user_name;
            state.admin_flag = user.admin_flag;
            state.mail_address = user.mail_address;
        }
    },
    actions: {
        auth({commit}, user) {
            commit('updateUser', user);
        }
    },
    modules: {},
    getters: {
        getToken(state) {
            return state.token
        }
    }
})

export default store