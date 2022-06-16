<template>
    <div>
        <HeaderComponent />
        <p>{{results[0].Time}}</p>
        <div class="statisticsTable">
            <table>
                <thead>
                    <tr>
                        <th id="rank">Rank</th>
                        <th id="number">Number</th>
                        <th id="rate">Rate</th>
                        <th id="count">Count</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(row, index) in results" :key="row.Number">
                        <td>{{index + 1}}</td>
                        <td>{{row.Number}}</td>
                        <td>{{row.Rate}}</td>
                        <td>{{row.Count}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'Loto6StatisticsPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            results:""
        }
    },
    mounted() {
        this.getLoto6Statistics()
    },
    methods: {
        async getLoto6Statistics() {
            const results = await axios.get("/getLoto6Statistics");
            this.results = results.data
        }
    }
};
</script>

<style>
.statisticsTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.statisticsTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
</style>
