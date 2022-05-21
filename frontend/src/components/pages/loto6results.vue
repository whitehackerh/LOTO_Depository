<template>
  <div>
    <HeaderComponent />
    <p>Results: {{results}}</p>
  </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'Loto6ResultPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            results:""
        }
    },
    mounted() {
        this.getLoto6Results()
    },
    methods: {
        async getLoto6Results() {
            const results = await axios.get("/getLoto6Results");
            this.results = results.data
        }
    }
};
</script>