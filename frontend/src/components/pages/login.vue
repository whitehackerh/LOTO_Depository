<template>
    <div>
    <HeaderComponent />
        <h2>テストページ2</h2>
  </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
export default {
    name: 'LoginPage',
    components: {
        HeaderComponent
    }
};
</script>