<template>
    <HeaderComponent />
    <div class="inputTable">
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="7">Input Numbers</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>{{time}}</td>
                    <td><input type="text" v-model="input_number_1"></td>
                    <td><input type="text" v-model="input_number_2"></td>
                    <td><input type="text" v-model="input_number_3"></td>
                    <td><input type="text" v-model="input_number_4"></td>
                    <td><input type="text" v-model="input_number_5"></td>
                    <td><input type="text" v-model="input_number_6"></td>
                    <td><input type="text" v-model="input_number_7"></td>
                </tr>
            </tbody>
        </table>
    </div>
    <br><br>
    <button @click="setLoto7Results()">Register</button>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'RegisterLoto7ResultsPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            time: '',
            input_number_1: '',
            input_number_2: '',
            input_number_3: '',
            input_number_4: '',
            input_number_5: '',
            input_number_6: '',
            input_number_7: '',
        }
    },
    mounted() {
        this.getNewestLoto7Result()
    },
    methods: {
        async setLoto7Results() {
            await axios.post("/setLoto7Results", {body: {time: this.time, input_number_1: this.input_number_1, input_number_2: this.input_number_2, input_number_3: this.input_number_3,
                 input_number_4: this.input_number_4, input_number_5: this.input_number_5, input_number_6: this.input_number_6, input_number_7: this.input_number_7}})
            .then(function (response) {
                console.log(response);
            });
        },
        async getNewestLoto7Result() {
            const result = await axios.get("/getNewestLoto7Result");
            this.time = String(result.data[0].Time)
        }
    }
}
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