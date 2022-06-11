<template>
  <div>
    <HeaderComponent />
    <div class="resultTable">
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="6">Numbers</th>
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