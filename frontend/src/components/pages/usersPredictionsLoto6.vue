<template>
    <div>
    <HeaderComponent />
    <button id="buttonDownloadLoto6UsersPredictions" v-show="predictions[0].existBeforeResultAnnouncement==true || predictions[0].existAfterResultAnnouncement == true" @click="downloadLoto6UsersPredictions()">Download CSV</button><br><br>
    <div class="beforeResultAnnouncementTable" v-show="predictions[0].existBeforeResultAnnouncement==true">
        <h3>Before Result Announcement</h3>
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="6">Numbers</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row in predictions" v-show="row.beforeResultAnnouncement==true" :key="row.Time">
                    <td id="LinkToDetail" @click="movePredictionsDetailLoto6(row.Time)">{{row.Time}}</td>
                    <td>{{row.Number_1}}</td>
                    <td>{{row.Number_2}}</td>
                    <td>{{row.Number_3}}</td>
                    <td>{{row.Number_4}}</td>
                    <td>{{row.Number_5}}</td>
                    <td>{{row.Number_6}}</td>
                </tr>
            </tbody>
        </table>
    </div>
    <div class="usersPredictionsTable" v-show="predictions[0].existAfterResultAnnouncement == true">
        <h3 id="Records">Records : {{predictions[0].Records}}</h3>
        <h3 id="Average">Average : {{Math.floor(predictions[0].Average * 1000) / 1000}}</h3>
        <h3 id="Rate">Rate : {{Math.floor(predictions[0].Rate * 1000) / 1000}}%</h3><br>
        <table v-show="predictions[0].Time!=0">
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="6">Numbers</th>
                    <th id="match">Matches</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row in predictions" v-show="row.beforeResultAnnouncement==false" :key="row.Time">
                    <td id="LinkToDetail" @click="movePredictionsDetailLoto6(row.Time)">{{row.Time}}</td>
                    <td>{{row.Number_1}}</td>
                    <td>{{row.Number_2}}</td>
                    <td>{{row.Number_3}}</td>
                    <td>{{row.Number_4}}</td>
                    <td>{{row.Number_5}}</td>
                    <td>{{row.Number_6}}</td>
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
    name: 'UsersPredictionsLoto6Page',
    components: {
        HeaderComponent
    },
    data() {
        return {
            predictions: '',
        }
    },
    mounted() {
        this.moveUserPredictionsLoto6()
        this.getLoto6UsersPredicitions()
    },
    methods: {
        moveUserPredictionsLoto6() {
            this.$router.push({
                name: 'usersPredictionsLoto6',
                params: { id: this.$store.getters.getUserId}
            })
        },
        async getLoto6UsersPredicitions() {
            const predictions = await axios.post("/getLoto6UsersPredictions", {body: {user_id: this.$store.getters.getUserId}});
            this.predictions = predictions.data
        },
        movePredictionsDetailLoto6(time) {
            this.$router.push({
                name: 'usersPredictionsDetailLoto6',
                params: { id: this.$store.getters.getUserId, time: time}
            })
        },
        async downloadLoto6UsersPredictions() {
            axios.post("/downloadLoto6UsersPredictions", {body: {user_id: this.$store.getters.getUserId}}, { responseType: 'blob'})
            .then((res) => {
                const bom = new Uint8Array([0xef, 0xbb, 0xbf]);
                const url = window.URL.createObjectURL(new Blob([bom, res.data]));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', 'Loto6Predictions.csv');
                document.body.appendChild(link);
                link.click();
                window.URL.revokeObjectURL(url);
            });
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