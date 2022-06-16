<template>
    <div>
        <HeaderComponent />
        <p>Time : {{results[0].Time}}</p><br><br>
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
                        <td>{{Math.floor(row.Rate * 1000) / 1000}}%</td>
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
    name: 'Loto7StatisticsPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            results:""
        }
    },
    mounted() {
        this.getLoto7Statistics()
    },
    methods: {
        async getLoto7Statistics() {
            const results = await axios.get("/getLoto7Statistics");
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
p {
    float: left;
}
</style>
