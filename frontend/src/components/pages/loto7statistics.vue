<template>
    <div>
        <HeaderComponent />
        <button @click="downloadLoto7Statistics()">Download CSV</button><br><br>
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
                    <tr v-for="row in results" :key="row.Number">
                        <td>{{row.Rank}}</td>
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
        },
        async downloadLoto7Statistics() {
            axios.get("/downloadLoto7Statistics", { responseType: 'blob'})
                .then((res) => {
                    const url = window.URL.createObjectURL(new Blob([res.data]))
                    const link = document.createElement('a')
                    link.href = url
                    link.setAttribute('download', 'Loto7Statistics.csv')
                    document.body.appendChild(link)
                    link.click()
                    window.URL.revokeObjectURL(url)
                })
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
