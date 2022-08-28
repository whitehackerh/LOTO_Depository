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
            <button @click="editLoto6UsersPredictions()" :disabled="isDisabled">Register</button>
            <button @click="deleteLoto6UsersPredictions()" :disabled="isDisabled">Delete</button>
            <table>
                <thead>
                    <tr>
                        <th><input type="checkbox" v-model="checkAll" @click="checkAllMethod()" id="checkAll"></th>
                        <th colspan="6">Numbers</th>
                        <th v-show="predictions[0].beforeResultAnnouncement==false">Matches</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(prediction, index) in predictions" :key="prediction.Time_Id">
                        <td><input type="checkbox" :value="prediction.Time_Id" v-model="checkSingle[index]" @click="clickCheckbox(prediction.Time_Id, index)"></td>
                        <input type="hidden" :value="predictions[index].Time" v-model="predictions[index].Time">
                        <input type="hidden" :value="predictions[index].Time_Id" v-model="predictions[index].Time_Id">
                        <td><input type="text" v-model="predictions[index].Number_1" :readonly="read[index]"></td>
                        <td><input type="text" v-model="predictions[index].Number_2" :readonly="read[index]"></td>
                        <td><input type="text" v-model="predictions[index].Number_3" :readonly="read[index]"></td>
                        <td><input type="text" v-model="predictions[index].Number_4" :readonly="read[index]"></td>
                        <td><input type="text" v-model="predictions[index].Number_5" :readonly="read[index]"></td>
                        <td><input type="text" v-model="predictions[index].Number_6" :readonly="read[index]"></td>
                        <td v-show="predictions[0].beforeResultAnnouncement==false">{{predictions[index].Matches}}</td>
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
            result: [],
            predictions: [],
            indexPredictions: [],
            checked: [],
            checkAll: false,
            checkSingle: [],
            read: [],
            isDisabled: true,
        }
    },
    mounted() {
        this.getLoto6UsersPredictionsDetail()
    },
    methods: {
        async getLoto6UsersPredictionsDetail() {
            const response = await axios.post("/getLoto6UsersPredictionsDetail", {body: {user_id: this.$store.getters.getUserId, time: this.$route.params.time}})
            this.result = response.data.Result;
            this.predictions = response.data.Predictions;
            this.indexPredictions = this.setIndex(this.predictions);
            this.read = [];
            for (let i = 0; i < Object.keys(this.predictions).length; i++) {
                this.read[i] = true;
                this.checkSingle.push(false);
            }
        },
        setIndex(predictions) {
            for (let i = 0; i < Object.keys(predictions).length; i++) {
                this.indexPredictions.push(i);
            }
        },
        clickCheckbox(time_id, index) {
            if (this.checkSingle[index]) {
                this.checkSingle[index] = false;
                this.read[index] = true;
            } else {
                this.checkSingle[index] = true;
                this.read[index] = false;
            }
            if (this.checked.includes(time_id)) {
                let index = this.checked.indexOf(time_id);
                this.checked.splice(index, 1);
                this.checkAll = false;
                if (this.checked.length == 0) {
                    this.isDisabled = true;
                }
                return;
            } else {
                this.isDisabled = false;
                this.checked.push(time_id);
                if (this.checked.length == Object.keys(this.predictions).length) {
                    this.checkAll = true;
                    return;
                } else {
                    this.checkAll = false;
                    return;
                }
            }
        },
        checkAllMethod() {
            if (this.checkAll) {
                this.checkAll = false;
                this.isDisabled = true;
                this.checked = [];
                this.checkSingle = [];
                this.read = [];
                for (let i = 0; i < Object.keys(this.predictions).length; i++) {
                    this.checkSingle.push(false);
                    this.read.push(true);
                }
                return;
            } else {
                this.checkAll = true;
                this.isDisabled = false;
                this.checked = [];
                this.checkSingle = [];
                this.read = [];
                for (let i = 0; i < Object.keys(this.predictions).length; i++) {
                    this.checked.push(this.predictions[i].Time_Id);
                    this.checkSingle.push(true);
                    this.read.push(false);
                }
                return;
            }
        },
        async editLoto6UsersPredictions() {
            let parameter = [];
            for (let i = 0; i < Object.keys(this.predictions).length; i++) {
                if (this.checkSingle[i]) {
                    this.predictions[i].Time = typeof this.predictions[i].Time == 'string' ? this.predictions[i].Time : String(this.predictions[i].Time);
                    this.predictions[i].Time_Id = typeof this.predictions[i].Time_Id == 'string' ? this.predictions[i].Time_Id : String(this.predictions[i].Time_Id);
                    this.predictions[i].Number_1 = typeof this.predictions[i].Number_1 == 'string' ? this.predictions[i].Number_1 : String(this.predictions[i].Number_1);
                    this.predictions[i].Number_2 = typeof this.predictions[i].Number_2 == 'string' ? this.predictions[i].Number_2 : String(this.predictions[i].Number_2);
                    this.predictions[i].Number_3 = typeof this.predictions[i].Number_3 == 'string' ? this.predictions[i].Number_3 : String(this.predictions[i].Number_3);
                    this.predictions[i].Number_4 = typeof this.predictions[i].Number_4 == 'string' ? this.predictions[i].Number_4 : String(this.predictions[i].Number_4);
                    this.predictions[i].Number_5 = typeof this.predictions[i].Number_5 == 'string' ? this.predictions[i].Number_5 : String(this.predictions[i].Number_5);
                    this.predictions[i].Number_6 = typeof this.predictions[i].Number_6 == 'string' ? this.predictions[i].Number_6 : String(this.predictions[i].Number_6);
                    this.predictions[i].Matches = typeof this.predictions[i].Matches == 'string' ? this.predictions[i].Mathces : String(this.predictions[i].Matches);
                    this.predictions[i].User_id = typeof this.predictions[i].User_Id == 'string' ? this.predictions[i].User_Id : String(this.predictions[i].User_Id);
                    if (i == 0) {
                        this.predictions[i].Average = typeof this.predictions[i].Average == 'string' ? this.predictions[i].Average : String(this.predictions[i].Average);
                        this.predictions[i].Rate = typeof this.predictions[i].Rate == 'string' ? this.predictions[i].Rate : String(this.predictions[i].Rate);
                        this.predictions[i].Records = typeof this.predictions[i].Records == 'string' ? this.predictions[i].Records : String(this.predictions[i].Records);
                        this.predictions[i].Total = typeof this.predictions[i].Total == 'string' ? this.predictions[i].Total : String(this.predictions[i].Total);
                        this.predictions[i].beforeResultAnnouncement = typeof this.predictions[i].beforeResutAnnouncement == 'boolean' ? this.predictions[i].beforeResultAnnouncement : this.predictions[i].beforeResultAnnouncement.toString();
                    }
                    parameter.push(this.predictions[i]);
                }
            }
            await axios.post("/editLoto6UsersPredictionsDetail", {body: {user_id: this.$store.getters.getUserId, predictions: parameter}});
        },
        async deleteLoto6UsersPredictions() {
            let parameter = [];
            for (let i = 0; i < Object.keys(this.predictions).length; i++) {
                if (this.checkSingle[i]) {
                    this.predictions[i].Time = typeof this.predictions[i].Time == 'string' ? this.predictions[i].Time : String(this.predictions[i].Time);
                    this.predictions[i].Time_Id = typeof this.predictions[i].Time_Id == 'string' ? this.predictions[i].Time_Id : String(this.predictions[i].Time_Id);
                    this.predictions[i].Number_1 = typeof this.predictions[i].Number_1 == 'string' ? this.predictions[i].Number_1 : String(this.predictions[i].Number_1);
                    this.predictions[i].Number_2 = typeof this.predictions[i].Number_2 == 'string' ? this.predictions[i].Number_2 : String(this.predictions[i].Number_2);
                    this.predictions[i].Number_3 = typeof this.predictions[i].Number_3 == 'string' ? this.predictions[i].Number_3 : String(this.predictions[i].Number_3);
                    this.predictions[i].Number_4 = typeof this.predictions[i].Number_4 == 'string' ? this.predictions[i].Number_4 : String(this.predictions[i].Number_4);
                    this.predictions[i].Number_5 = typeof this.predictions[i].Number_5 == 'string' ? this.predictions[i].Number_5 : String(this.predictions[i].Number_5);
                    this.predictions[i].Number_6 = typeof this.predictions[i].Number_6 == 'string' ? this.predictions[i].Number_6 : String(this.predictions[i].Number_6);
                    this.predictions[i].Matches = typeof this.predictions[i].Matches == 'string' ? this.predictions[i].Mathces : String(this.predictions[i].Matches);
                    this.predictions[i].User_id = typeof this.predictions[i].User_Id == 'string' ? this.predictions[i].User_Id : String(this.predictions[i].User_Id);
                    if (i == 0) {
                        this.predictions[i].Average = typeof this.predictions[i].Average == 'string' ? this.predictions[i].Average : String(this.predictions[i].Average);
                        this.predictions[i].Rate = typeof this.predictions[i].Rate == 'string' ? this.predictions[i].Rate : String(this.predictions[i].Rate);
                        this.predictions[i].Records = typeof this.predictions[i].Records == 'string' ? this.predictions[i].Records : String(this.predictions[i].Records);
                        this.predictions[i].Total = typeof this.predictions[i].Total == 'string' ? this.predictions[i].Total : String(this.predictions[i].Total);
                        this.predictions[i].beforeResultAnnouncement = typeof this.predictions[i].beforeResutAnnouncement == 'boolean' ? this.predictions[i].beforeResultAnnouncement : this.predictions[i].beforeResultAnnouncement.toString();
                    }
                    parameter.push(this.predictions[i]);
                }
            }
            await axios.post("/deleteLoto6UsersPredictionsDetail", {body: {user_id: this.$store.getters.getUserId, predictions: parameter}});
        }
    }
}
</script>

<style>
.resultTable {
    text-align: center;
}
.resultTable th, td{
    width: 100px;
}
.resultTable th {
    border: 1px solid #000066;
    background: #52a3fe;
}
.resultTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
.prefictionsTable {
    text-align: center;
}
.predictionsTable th, td{
    width: 100px;
}
.predictionsTable th {
    border: 1px solid #000066;
    background: #52a3fe;
}
.predictionsTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
h3 {
    color: white;
}
</style>     
