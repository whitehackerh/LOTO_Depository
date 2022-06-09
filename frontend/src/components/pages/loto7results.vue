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
    name: 'Loto7ResultsPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            results:""
        }
    },
    mounted() {
        this.getLoto7Results()
    },
    methods: {
        async getLoto7Results() {
            const results = await axios.get("/getLoto7Results");
            this.results = results.data
        }
    }
};
</script>