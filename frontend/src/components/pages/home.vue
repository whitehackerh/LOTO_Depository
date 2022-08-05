<template>
    <div id="app">
        <HeaderComponent />
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
export default {
    name: 'HomePage',
    components: {
        HeaderComponent
    }
};
</script>
<style>
#app {
  font-family: "Century Gothic", Helvetica, Arial, sans-serif;
  font-weight:800; 
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /*background-color: #004CFF;*/
  /* margin-top: 60px; */
}
body {
    background-image: url("~@/assets/images/earth.jpg");
  background-repeat: repeat;
}
</style>