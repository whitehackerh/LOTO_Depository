<template>
    <div>
    <HeaderComponent />
        <div class="beforeResultAnnouncementTable" v-show="predictions[0].Time!=0 && predictions[0].existBeforeResultAnnouncement==true">
        <h3>Before Result Announcement</h3>
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="7">Numbers</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row in predictions" v-show="row.beforeResultAnnouncement==true" :key="row.Time">
                    <td id="LinkToDetail" @click="movePredictionsDetailLoto7(row.Time)">{{row.Time}}</td>
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
    <div class="usersPredictionsTable" v-show="predictions[0].Time!=0">
        <h3 id="Records">Records : {{predictions[0].Records}}</h3>
        <h3 id="Average">Average : {{Math.floor(predictions[0].Average * 1000) / 1000}}</h3>
        <h3 id="Rate">Rate : {{Math.floor(predictions[0].Rate * 1000) / 1000}}%</h3><br>
        <table v-show="predictions[0].Time!=0">
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="7">Numbers</th>
                    <th id="match">Matches</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row in predictions" v-show="row.beforeResultAnnouncement==false" :key="row.Time">
                    <td id="LinkToDetail" @click="movePredictionsDetailLoto7(row.Time)">{{row.Time}}</td>
                    <td>{{row.Number_1}}</td>
                    <td>{{row.Number_2}}</td>
                    <td>{{row.Number_3}}</td>
                    <td>{{row.Number_4}}</td>
                    <td>{{row.Number_5}}</td>
                    <td>{{row.Number_6}}</td>
                    <td>{{row.Number_7}}</td>
                    <td>{{row.Matches}}</td>
                </tr>
            </tbody>
        </table>
    </div>
    <h3 v-show="predictions[0].Time==0">No Data</h3>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';

export default {
    name: 'UsersPredictionsLoto7Page',
    components: {
        HeaderComponent
    },
    data() {
        return {
            predictions:""
        }
    },
    mounted() {
        this.moveUserPredictionsLoto7()
        this.getLoto7UsersPredicitions()
    },
    methods: {
        moveUserPredictionsLoto7() {
            this.$router.push({
                name: 'usersPredictionsLoto7',
                params: { id: this.$store.getters.getUserId}
            })
        },
        async getLoto7UsersPredicitions() {
            const predictions = await axios.post("/getLoto7UsersPredictions", {body: {user_id: this.$store.getters.getUserId}});
            this.predictions = predictions.data
        },
        movePredictionsDetailLoto7(time) {
            this.$router.push({
                name: 'usersPredictionsDetailLoto7',
                params: { id: this.$store.getters.getUserId, time: time}
            })
        }
    }
}
</script>

<style>
.usersPredictionsTable th,td {
    width: 100px;
}
.usersPredictionsTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.usersPredictionsTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
.beforeResultAnnouncementTable th,td {
    width: 100px;
}
.beforeResultAnnouncementTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.beforeResultAnnouncementTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
#LinkToDetail {
    color: blue;
    text-decoration: underline;
}
</style>