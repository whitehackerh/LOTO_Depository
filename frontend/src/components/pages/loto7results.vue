<template>
    <div>
        <HeaderComponent />
        <button @click="downloadLoto7Results()">Download CSV</button><br><br>
        <div class="resultTable">
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="7">Numbers</th>
                    <th>Edit</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row in results" :key="row.Time">
                    <td>{{row.Time}}</td>
                    <td>{{row.Number_1}}</td>
                    <td>{{row.Number_2}}</td>
                    <td>{{row.Number_3}}</td>
                    <td>{{row.Number_4}}</td>
                    <td>{{row.Number_5}}</td>
                    <td>{{row.Number_6}}</td>
                    <td>{{row.Number_7}}</td>
                    <td><button @click="moveEditLoto7Result(row.Time)">Edit</button></td>
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
        },
        async downloadLoto7Results() {
            axios.get("/downloadLoto7Results", { responseType: 'blob',})
                .then((res) => {
                    const url = window.URL.createObjectURL(new Blob([res.data]))
                    const link = document.createElement('a')
                    link.href = url
                    link.setAttribute('download', 'Loto7Results.csv')
                    document.body.appendChild(link)
                    link.click()
                    window.URL.revokeObjectURL(url)
                })
        },
        moveEditLoto7Result(time) {
            this.$router.push({
                name: 'editLoto7Result',
                params: { id: time}
            })
        }
    }
};
</script>

<style>
.resultTable th,td {
    width: 100px;
}
.resultTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.resultTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
</style>