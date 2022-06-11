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
                    <td><input type="text" v-model="data_1"></td>
                    <td><input type="text" v-model="data_2"></td>
                    <td><input type="text" v-model="data_3"></td>
                    <td><input type="text" v-model="data_4"></td>
                    <td><input type="text" v-model="data_5"></td>
                    <td><input type="text" v-model="data_6"></td>
                    <td><button @click="(determineLoto6Expectation())">Check</button></td>
                    <td><div v-show="infoflag">{{info.message}}</div></td>
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
            data_1: '',
            data_2: '',
            data_3: '',
            data_4: '',
            data_5: '',
            data_6: '',
            info: '',
            infoflag: false
        }
    },
    methods: {
        async determineLoto6Expectation() {
            const results = await axios.post("determineLoto6Expectation", {body: {time: this.time, data_1: this.data_1, data_2: this.data_2, data_3: this.data_3,
                 data_4: this.data_4, data_5: this.data_5, data_6: this.data_6}})
            this.infoflag = true
            this.info = results.data
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