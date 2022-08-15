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
                <tr v-for="(prediction, index) in predictions" :key="index">
                    <td><input type="text" v-model="predictions[index].time"></td>
                    <td><input type="text" v-model="predictions[index].input_number_1"></td>
                    <td><input type="text" v-model="predictions[index].input_number_2"></td>
                    <td><input type="text" v-model="predictions[index].input_number_3"></td>
                    <td><input type="text" v-model="predictions[index].input_number_4"></td>
                    <td><input type="text" v-model="predictions[index].input_number_5"></td>
                    <td><input type="text" v-model="predictions[index].input_number_6"></td>
                    <td><button @click="(determineLoto6Prediction(index))">Check</button></td>
                    <td><div v-show="predictions[index].infoflag=='1'">{{predictions[index].info}}</div></td>
                </tr>
            </tbody>
        </table>
    </div>
    <br><br>
    <button @click="setLoto6Predictions()">Register</button>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'PredictLoto6Page',
    components: {
        HeaderComponent
    },
    data() {
        return {
            predictions: [{
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
        async determineLoto6Prediction(index) {
            const results = await axios.post("/determineLoto6Prediction", {body: {time: this.predictions[index].time, input_number_1: this.predictions[index].input_number_1, input_number_2: this.predictions[index].input_number_2, input_number_3: this.predictions[index].input_number_3,
                 input_number_4: this.predictions[index].input_number_4, input_number_5: this.predictions[index].input_number_5, input_number_6: this.predictions[index].input_number_6}})
            this.predictions[index].infoflag = '1'
            this.predictions[index].info = results.data[0].Result
        },
        addRow() {
            this.predictions.push({
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
        async setLoto6Predictions() {
            await axios.post("/setLoto6Predictions", {body: {user_id: this.$store.getters.getUserId, predictions: this.predictions}})
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