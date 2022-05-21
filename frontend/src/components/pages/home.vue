<template>
    <div id="app">
        <HeaderComponent />
        <h2>テストページ1</h2>
        <router-link to="/login">次のページ</router-link>
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
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
</style>