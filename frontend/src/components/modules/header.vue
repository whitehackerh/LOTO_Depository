<template #header>
  <b-nav tabs>
    <b-nav-item><router-link to="/"><img id="logo" src="../../assets/images/toplogo_w_trans.png"></router-link></b-nav-item>
    <b-nav-item-dropdown id="loto6list" text="LOTO6" toggle-class="nav-link-custom">
      <b-dropdown-item><router-link to="/loto6results">Result</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/loto6statistics">Statistics</router-link></b-dropdown-item>
      <b-dropdown-item v-show="this.$store.getters.getAuthority=='1'"><router-link to="/registerloto6results">Register</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/predictloto6">Predict</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/usersPredictionsLoto6/:id">Predictions</router-link></b-dropdown-item>
    </b-nav-item-dropdown>
    <b-nav-item-dropdown id="loto7list" text="LOTO7" toggle-class="nav-link-custom">
      <b-dropdown-item><router-link to="/loto7results">Result</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/loto7statistics">Statistics</router-link></b-dropdown-item>
      <b-dropdown-item v-show="this.$store.getters.getAuthority=='1'"><router-link to="/registerloto7results">Register</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/predictloto7">Predict</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/usersPredictionsLoto7/:id">Predictions</router-link></b-dropdown-item>
    </b-nav-item-dropdown>
    <b-nav-item-dropdown id="user" text="User" toggle-class="nav-link-custom">
      <b-dropdown-item v-show="this.$store.getters.getUserId==''"><router-link to="/login">Login</router-link></b-dropdown-item>
      <b-dropdown-item v-show="this.$store.getters.getUserId!=''" @click="Logout()">Logout</b-dropdown-item>
    </b-nav-item-dropdown>
  </b-nav>
</template>
<script>
export default {
    name: 'HeaderComponent',
    methods: {
      async Logout() {
        this.$store.dispatch("auth", {
            user_id: '',
            user_name:'',
            mail_address: '',
            admin_flag: '', 
            token: '',
        })
        this.$router.push('/');
      },
    }
};
</script>
<style>
  #header {
    position: sticky;
    background: white;
    opacity: 0.3;
  }
  #logo {
    width: auto;
    height: 75px;
  }
  #loto6list {
    
  }
  #loto7list {

  }
</style>