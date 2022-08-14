<template>
    <HeaderComponent />
    <div class="addRow">
        <button @click="addRow()">Add Row</button>
    </div>
    <div class="inputTable">
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="6">Input Numbers</th>
                    <th colspan="2">Check</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(expectation, index) in expectations" :key="index">
                    <td><input type="text" v-model="expectations[index].time"></td>
                    <td><input type="text" v-model="expectations[index].input_number_1"></td>
                    <td><input type="text" v-model="expectations[index].input_number_2"></td>
                    <td><input type="text" v-model="expectations[index].input_number_3"></td>
                    <td><input type="text" v-model="expectations[index].input_number_4"></td>
                    <td><input type="text" v-model="expectations[index].input_number_5"></td>
                    <td><input type="text" v-model="expectations[index].input_number_6"></td>
                    <td><button @click="(determineLoto6Expectation(index))">Check</button></td>
                    <td><div v-show="expectations[index].infoflag=='1'">{{expectations[index].info}}</div></td>
                </tr>
            </tbody>
        </table>
    </div>
    <br><br>
    <button @click="setLoto6Expectations()">Register</button>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'ExpectLoto6Page',
    components: {
        HeaderComponent
    },
    data() {
        return {
            expectations: [{
                time: '',
                input_number_1: '',
                input_number_2: '',
                input_number_3: '',
                input_number_4: '',
                input_number_5: '',
                input_number_6: '',
                info: '',
                infoflag: '2'
            }]
        }
    },
    methods: {
        async determineLoto6Expectation(index) {
            const results = await axios.post("/determineLoto6Expectation", {body: {time: this.expectations[index].time, input_number_1: this.expectations[index].input_number_1, input_number_2: this.expectations[index].input_number_2, input_number_3: this.expectations[index].input_number_3,
                 input_number_4: this.expectations[index].input_number_4, input_number_5: this.expectations[index].input_number_5, input_number_6: this.expectations[index].input_number_6}})
            this.expectations[index].infoflag = '1'
            this.expectations[index].info = results.data[0].Result
        },
        addRow() {
            this.expectations.push({
                time: '',
                input_number_1: '',
                input_number_2: '',
                input_number_3: '',
                input_number_4: '',
                input_number_5: '',
                input_number_6: '',
                info: '',
                infoflag: '2'
            })
        },
        async setLoto6Expectations() {
            await axios.post("/setLoto6Expectations", {body: {user_id: this.$store.getters.getUserId, expectations: this.expectations}})
            .then(function(response) {
                console.log(response);
            })
        }
    }
}
</script>

<style>
.inputTable th,td {
    width: 60px;
}
.inputTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.inputTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
.addRow {
    text-align: left;
}
</style>