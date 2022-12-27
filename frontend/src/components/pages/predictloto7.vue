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
                    <th colspan="7">Input Numbers</th>
                    <th>Random</th>
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
                    <td><input type="text" v-model="predictions[index].input_number_7"></td>
                    <td><button @click="(generateRandomNumbers(index))">Execute</button></td>
                    <td><button @click="(determineLoto7Prediction(index))">Check</button></td>
                    <td><div v-show="predictions[index].infoflag=='1'">{{predictions[index].info}}</div></td>
                </tr>
            </tbody>
        </table>
    </div>
    <br><br>
    <button @click="setLoto7Predictions()">Register</button>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
import {getRandomInt} from '../../util/functionUtil.js';
export default {
    name: 'PredictLoto7Page',
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
                input_number_7: '',
                info: '',
                infoflag: '2'
            }]
        }
    },
    methods: {
        async determineLoto7Prediction(index) {
            const results = await axios.post("/determineLoto7Prediction", {body: {time: this.predictions[index].time, input_number_1: this.predictions[index].input_number_1, input_number_2: this.predictions[index].input_number_2, input_number_3: this.predictions[index].input_number_3,
                 input_number_4: this.predictions[index].input_number_4, input_number_5: this.predictions[index].input_number_5, input_number_6: this.predictions[index].input_number_6, input_number_7: this.predictions[index].input_number_7}})
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
                input_number_7: '',
                info: '',
                infoflag: '2'
            })
        },
        async setLoto7Predictions() {
            await axios.post("/setLoto7Predictions", {body: {user_id: this.$store.getters.getUserId, predictions: this.predictions}})
            .then(function(response) {
                console.log(response);
            })
        },
        generateRandomNumbers(index) {
            let randomNumbers = [];
            let isDuplication = true;
            let pushCount = 0;
            for (; pushCount < 7; pushCount++) {
                while (isDuplication) {
                    const number = getRandomInt(1, 37);
                    if (randomNumbers.includes(number)) {
                        continue;
                    }
                    randomNumbers.push(number);
                    isDuplication = false;
                }
                isDuplication = true;
            }
            randomNumbers.sort((a, b) => a - b);
            this.predictions[index].infoflag = '2';
            this.predictions[index].info = '';
            for (let i = 0; i < 7; i++) {
                randomNumbers[i] = String(randomNumbers[i]);
                switch (i) {
                    case 0:
                        this.predictions[index].input_number_1 = randomNumbers[i];
                        break;
                    case 1:
                        this.predictions[index].input_number_2 = randomNumbers[i];
                        break;
                    case 2:
                        this.predictions[index].input_number_3 = randomNumbers[i];
                        break;
                    case 3:
                        this.predictions[index].input_number_4 = randomNumbers[i];
                        break;
                    case 4:
                        this.predictions[index].input_number_5 = randomNumbers[i];
                        break;
                    case 5:
                        this.predictions[index].input_number_6 = randomNumbers[i];
                        break;
                    case 6:
                        this.predictions[index].input_number_7 = randomNumbers[i];
                        break;
                    default:
                        break;
                }
            }
        }
    }
}
</script>

<style>
.inputTable th {
    border: 1px solid #000066;
    background: #66CCFF;
    width: 100px;
}
.inputTable td {
    border: 1px solid #000066;
    background: #ffffff;
    width: 100px;
}
.addRow {
    text-align: left;
}
</style>