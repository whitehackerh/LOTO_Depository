<template>
    <div>
        <HeaderComponent />
        <h3>Time : {{this.$route.params.time}}</h3>
        <div class="resultTable">
            <table>
                <thead>
                    <tr>
                        <th colspan="7">Numbers</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="row in result" :key="row.Time">
                        <td>{{row.Number_1}}</td>
                        <td>{{row.Number_2}}</td>
                        <td>{{row.Number_3}}</td>
                        <td>{{row.Number_4}}</td>
                        <td>{{row.Number_5}}</td>
                        <td>{{row.Number_6}}</td>
                        <td>{{row.Number_7}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
        <br><br>
        <h3 id="Records">Records : {{predictions[0].Records}}</h3>
        <h3 id="Average">Average : {{Math.floor(predictions[0].Average * 1000) / 1000}}</h3>
        <h3 id="Rate">Rate : {{Math.floor(predictions[0].Rate * 1000) / 1000}}%</h3><br>
        <div class="predictionsTable">
            <table>
                <thead>
                    <tr>
                        <th><input type="checkbox" @click="allSelect()" v-model="allSelected"></th>
                        <th colspan="7">Numbers</th>
                        <th>Matches</th>
                        <th>Edit</th>
                        <th>Register</th>
                        <th>Delete</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="prediction in predictions" :key="prediction.Time_Id">
                        <td><input type="checkbox" :value="prediction" v-model="selected"></td>
                        <td>{{prediction.Number_1}}</td>
                        <td>{{prediction.Number_2}}</td>
                        <td>{{prediction.Number_3}}</td>
                        <td>{{prediction.Number_4}}</td>
                        <td>{{prediction.Number_5}}</td>
                        <td>{{prediction.Number_6}}</td>
                        <td>{{prediction.Number_7}}</td>
                        <td>{{prediction.Matches}}</td>
                        <td><button @click="edit()">Edit</button></td>
                        <td><button @click="editLoto7UsersPrediction(prediction.Time_Id)">Register</button></td>
                        <td><button @click="deleteLoto7UsersPrediction(prediction.Time_Id)">Delete</button></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';

export default {
    name: 'UsersPredictionsDetailLoto7Page',
    components: {
        HeaderComponent
    },
    data() {
        return {
            result: '',
            predictions: '',
            allSelected: false,
            selected: [],
        }
    },
    mounted() {
        this.getLoto7UsersPredictionsDetail()
    },
    methods: {
        async getLoto7UsersPredictionsDetail() {
            const response = await axios.post("/getLoto7UsersPredictionsDetail", {body: {user_id: this.$store.getters.getUserId, time: this.$route.params.time}})
            this.result = response.data.Result
            this.predictions = response.data.Predictions
        }
    }
}
</script>

<style>
.resultTable th, td{
    width: 100px;
}
.resultTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.resultTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
.predictionsTable th, td{
    width: 100px;
}
.predictionsTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.predictionsTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
h3 {
    color: white;
}
</style>     
