<template>
    <HeaderComponent />
    <div class="inputTable">
        <table border="1">
            <thead>
                <tr>
                    <th style="width:120px;">Time</th>
                    <th style="width:100px;">①</th>
                    <th style="width:100px;">②</th>
                    <th style="width:100px;">③</th>
                    <th style="width:100px;">④</th>
                    <th style="width:100px;">⑤</th>
                    <th style="width:100px;">⑥</th>
                    <th style="width:100px;">⑦</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td><input type="text" style="width:120px;" v-model="time"></td>
                    <td><input type="text" style="width:100px;" v-model="data_1"></td>
                    <td><input type="text" style="width:100px;" v-model="data_2"></td>
                    <td><input type="text" style="width:100px;" v-model="data_3"></td>
                    <td><input type="text" style="width:100px;" v-model="data_4"></td>
                    <td><input type="text" style="width:100px;" v-model="data_5"></td>
                    <td><input type="text" style="width:100px;" v-model="data_6"></td>
                    <td><input type="text" style="width:100px;" v-model="data_7"></td>
                </tr>
            </tbody>
        </table>
    </div>
    <button @click="setLoto7Results()">register</button>
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
            data_1: '',
            data_2: '',
            data_3: '',
            data_4: '',
            data_5: '',
            data_6: '',
            data_7: '',
        }
    },
    methods: {
        async setLoto7Results() {
            await axios.post("/setLoto7Results", {body: {time: this.time, data_1: this.data_1, data_2: this.data_2, data_3: this.data_3,
                 data_4: this.data_4, data_5: this.data_5, data_6: this.data_6, data_7: this.data_7}})
            .then(function (response) {
                console.log(response);
            });
        }
    }
}
</script>