<template>
    <HeaderComponent />
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
                <tr>
                    <td><input type="text" v-model="time"></td>
                    <td><input type="text" v-model="input_number_1"></td>
                    <td><input type="text" v-model="input_number_2"></td>
                    <td><input type="text" v-model="input_number_3"></td>
                    <td><input type="text" v-model="input_number_4"></td>
                    <td><input type="text" v-model="input_number_5"></td>
                    <td><input type="text" v-model="input_number_6"></td>
                    <td><button @click="(determineLoto6Expectation())">Check</button></td>
                    <td><div v-show="infoflag">{{info}}</div></td>
                </tr>
            </tbody>
        </table>
    </div>
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
            time: '',
            input_number_1: '',
            input_number_2: '',
            input_number_3: '',
            input_number_4: '',
            input_number_5: '',
            input_number_6: '',
            info: '',
            infoflag: false
        }
    },
    methods: {
        async determineLoto6Expectation() {
            const results = await axios.post("determineLoto6Expectation", {body: {time: this.time, input_number_1: this.input_number_1, input_number_2: this.input_number_2, input_number_3: this.input_number_3,
                 input_number_4: this.input_number_4, input_number_5: this.input_number_5, input_number_6: this.input_number_6}})
            this.infoflag = true
            this.info = results.data[0].Result
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
</style>