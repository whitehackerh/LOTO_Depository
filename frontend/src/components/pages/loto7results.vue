<template>
    <div>
        <HeaderComponent />
        <div class="resultTable">
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th>①</th>
                    <th>②</th>
                    <th>③</th>
                    <th>④</th>
                    <th>⑤</th>
                    <th>⑥</th>
                    <th>⑦</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row in results" :key="row.Time">
                    <td>{{row.Time}}</td>
                    <td>{{row.Winning_number_1}}</td>
                    <td>{{row.Winning_number_2}}</td>
                    <td>{{row.Winning_number_3}}</td>
                    <td>{{row.Winning_number_4}}</td>
                    <td>{{row.Winning_number_5}}</td>
                    <td>{{row.Winning_number_6}}</td>
                    <td>{{row.Winning_number_7}}</td>
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