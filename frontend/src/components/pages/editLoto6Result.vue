<template>
    <div>
        <HeaderComponent />
        <h3>Time : {{this.$route.params.id}}</h3>
        <div class="editTable">
            <table>
                <thead>
                    <tr>
                        <th colspan="6">Numbers</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="row in detail" :key="row.Time">
                        <td>{{row.Number_1}}</td>
                        <td>{{row.Number_2}}</td>
                        <td>{{row.Number_3}}</td>
                        <td>{{row.Number_4}}</td>
                        <td>{{row.Number_5}}</td>
                        <td>{{row.Number_6}}</td>
                    </tr>
                    <tr>
                        <td><input type="text" v-model="input_number_1"></td>
                        <td><input type="text" v-model="input_number_2"></td>
                        <td><input type="text" v-model="input_number_3"></td>
                        <td><input type="text" v-model="input_number_4"></td>
                        <td><input type="text" v-model="input_number_5"></td>
                        <td><input type="text" v-model="input_number_6"></td>
                    </tr>
                </tbody>
            </table>
        </div>
        <br><br>
        <button  @click="editLoto6Result()">Register</button>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';

export default {
    name: 'EditLoto6ResultPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            detail: "",
            input_number_1: '',
            input_number_2: '',
            input_number_3: '',
            input_number_4: '',
            input_number_5: '',
            input_number_6: '',
        }
    },
    mounted() {
        this.getLoto6ResultDetail()
    },
    methods: {
        async getLoto6ResultDetail() {
            const result = await axios.post("/getLoto6ResultDetail", {body: {time: this.$route.params.id}})
            this.detail = result.data
        },
        async editLoto6Result() {
            await axios.post("/editLoto6Result", {body: {time: this.$route.params.id, input_number_1: this.input_number_1, input_number_2: this.input_number_2, input_number_3: this.input_number_3,
                 input_number_4: this.input_number_4, input_number_5: this.input_number_5, input_number_6: this.input_number_6}})
                 .then(function (response) {
                    console.log(response);
                 })
        }
    }
};
</script>

<style>
.editTable th,td {
    width: 100px;
}
.editTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.editTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
h3 {
    color: white;
}
</style>