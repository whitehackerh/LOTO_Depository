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
            // values: [
            //     {time: '', data_1: '', data_2: '', data_3: '',
            //      data_4: '', data_5: '', data_6: ''}
            // ]
        }
    },
    methods: {
        async setLoto6Results() {
            // const params = new URLSearchParams();
            // params.append('time', this.time);
            // params.append('data_1', this.data_1);
            // params.append('data_2', this.data_2);
            // params.append('data_3', this.data_3);
            // params.append('data_4', this.data_4);
            // params.append('data_5', this.data_5);
            // params.append('data_6', this.data_6);
            // await axios.post("/setLoto6Results", params);
            // let body = {'time': this.time, 'data_1': this.data_1, 'data_2': this.data_2, 'data_3': this.data_3,
            //     'data_4': this.data_4, 'data_5': this.data_5, 'data_6': this.data_6}
            // let requestBody = JSON.stringify(body);
            // await axios.post("/setLoto6Results", {'time': this.time, 'data_1': this.data_1, 'data_2': this.data_2, 'data_3': this.data_3,
            //     'data_4': this.data_4, 'data_5': this.data_5, 'data_6': this.data_6});
            // let a = {time: this.time, data_1: this.data_1, data_2: this.data_2, data_3: this.data_3,
            //     data_4: this.data_4, data_5: this.data_5, data_6: this.data_6};
            //     console.log(a);
            // let request = JSON.stringify(a);
            // await axios.post("/setLoto6Results", {time: this.time, data_1: this.data_1, data_2: this.data_2, data_3: this.data_3,
            //     data_4: this.data_4, data_5: this.data_5, data_6: this.data_6})
            //     .then(function (response) {
            //         // handle success
            //         console.log(response);
            //     });
            // await axios.post("/setLoto6Results", {time: 1, data_1: 1, data_2: 2, data_3: 3,
            //     data_4: 4, data_5: 5, data_6: 6})
            //     .then(function (response) {
            //         // handle success
            //         console.log(response);
            //     });
            await axios.post("/setLoto6Results", {body: {time: this.time, data_1: this.data_1, data_2: this.data_2, data_3: this.data_3,
                 data_4: this.data_4, data_5: this.data_5, data_6: this.data_6}})
                .then(function (response) {
                    // handle success
                    console.log(response);
                });
            //await axio
            //await axios.post("/setLoto6Results", request);
                // .then(function (response) {
                //     // handle success
                //     console.log(response);
                // });
        }
    }
};

// <td><input type="text" style="width:120px;" v-model="time"></td>
//                     <td><input type="text" style="width:100px;" v-model="data_1"></td>
//                     <td><input type="text" style="width:100px;" v-model="data_2"></td>
//                     <td><input type="text" style="width:100px;" v-model="data_3"></td>
//                     <td><input type="text" style="width:100px;" v-model="data_4"></td>
//                     <td><input type="text" style="width:100px;" v-model="data_5"></td>
//                     <td><input type="text" style="width:100px;" v-model="data_6"></td>
// <td><input type="text" style="width:120px;" v-model="values.time"></td>
//                     <td><input type="text" style="width:100px;" v-model="values.data_1"></td>
//                     <td><input type="text" style="width:100px;" v-model="values.data_2"></td>
//                     <td><input type="text" style="width:100px;" v-model="values.data_3"></td>
//                     <td><input type="text" style="width:100px;" v-model="values.data_4"></td>
//                     <td><input type="text" style="width:100px;" v-model="values.data_5"></td>
//                     <td><input type="text" style="width:100px;" v-model="values.data_6"></td>
</script>

