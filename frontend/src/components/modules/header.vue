<template #header>
  <b-nav tabs>
    <b-nav-item><router-link to="/"><img id="logo" src="../../assets/images/toplogo.png"></router-link></b-nav-item>
    <b-nav-item-dropdown id="loto6list" text="LOTO6" toggle-class="nav-link-custom">
      <b-dropdown-item><router-link to="/loto6results">Result</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/registerloto6results">Register</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/expectloto6">Expect</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/loto6statistics">Statistics</router-link></b-dropdown-item>
    </b-nav-item-dropdown>
    <b-nav-item-dropdown id="loto7list" text="LOTO7" toggle-class="nav-link-custom">
      <b-dropdown-item><router-link to="/loto7results">Result</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/registerloto7results">Register</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/expectloto7">Expect</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/loto7statistics">Statistics</router-link></b-dropdown-item>
    </b-nav-item-dropdown>
  </b-nav>
</template>
<script>
export default {
    name: 'HeaderComponent'
};
</script>
<style>
  #header {
    position: sticky;
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