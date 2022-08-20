<template>
    <div>
        <HeaderComponent />
        <h3>Time : {{this.$route.params.time}}</h3>
        <div class="resultTable">
            <table v-show="predictions[0].beforeResultAnnouncement==false">
                <thead>
                    <tr>
                        <th colspan="6">Numbers</th>
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
                    </tr>
                </tbody>
            </table>
            <h3 v-show="predictions[0].beforeResultAnnouncement==true">This Time is before the results are announced.</h3>
        </div>
        <br><br>
        <div class="predictionsTable" v-show="predictions[0].Time!=0">
            <h3 id="Records">Records : {{predictions[0].Records}}</h3>
            <h3 id="Average" v-show="predictions[0].beforeResultAnnouncement==false">Average : {{Math.floor(predictions[0].Average * 1000) / 1000}}</h3>
            <h3 id="Rate" v-show="predictions[0].beforeResultAnnouncement==false">Rate : {{Math.floor(predictions[0].Rate * 1000) / 1000}}%</h3><br>
            <table>
                <thead>
                    <tr>
                        <th><input type="checkbox" @click="allSelect()" v-model="allSelected"></th>
                        <th colspan="6">Numbers</th>
                        <th v-show="predictions[0].beforeResultAnnouncement==false">Matches</th>
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
                        <td v-show="predictions[0].beforeResultAnnouncement==false">{{prediction.Matches}}</td>
                        <td><button @click="edit()">Edit</button></td>
                        <td><button @click="editLoto6UsersPrediction(prediction.Time_Id)">Register</button></td>
                        <td><button @click="deleteLoto6UsersPrediction(prediction.Time_Id)">Delete</button></td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div>
            <h3 v-show="predictions[0].Time==0 && predictions[0].beforeResultAnnouncement == true">No Data</h3>
        </div>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';

export default {
    name: 'UsersPredictionsDetailLoto6Page',
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
        this.getLoto6UsersPredictionsDetail()
    },
    methods: {
        async getLoto6UsersPredictionsDetail() {
            const response = await axios.post("/getLoto6UsersPredictionsDetail", {body: {user_id: this.$store.getters.getUserId, time: this.$route.params.time}})
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
