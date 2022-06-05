<template #header>
  <b-nav tabs>
    <b-nav-item><router-link to="/"><img src="../../assets/images/toplogo.png"></router-link></b-nav-item>
    <b-nav-item-dropdown id="loto6list" text="LOTO6" toggle-class="nav-link-custom">
      <b-dropdown-item><router-link to="/loto6results">Result</router-link></b-dropdown-item>
      <b-dropdown-item><router-link to="/registerloto6results">Register</router-link></b-dropdown-item>
    </b-nav-item-dropdown>
    <b-nav-item>Disabled</b-nav-item>
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
</style>