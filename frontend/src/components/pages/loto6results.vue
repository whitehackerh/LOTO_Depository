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