<template>
    <HeaderComponent />
    <div class="inputTable">
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
                <tr>
                    <td><input type="text" v-model="time"></td>
                    <td><input type="text" v-model="data_1"></td>
                    <td><input type="text" v-model="data_2"></td>
                    <td><input type="text" v-model="data_3"></td>
                    <td><input type="text" v-model="data_4"></td>
                    <td><input type="text" v-model="data_5"></td>
                    <td><input type="text" v-model="data_6"></td>
                </tr>
            </tbody>
        </table>
    </div>
    <button @click="setLoto6Results()">register</button>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'RegisterLoto6ResultsPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            time: '',
            data_1: '',
            data_2: '',
            data_3: '',
            data_4: '',
            data_5: '',
            data_6: '',
        }
    },
    methods: {
        async setLoto6Results() {
            await axios.post("/setLoto6Results", {body: {time: this.time, data_1: this.data_1, data_2: this.data_2, data_3: this.data_3,
                 data_4: this.data_4, data_5: this.data_5, data_6: this.data_6}})
                .then(function (response) {
                    // handle success
                    console.log(response);
                });
        }
    }
};
</script>

<style>
.inputTable th,td {
    width: 100px;
}
.inputTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.inputTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
</style>