<template>
    <div>
        <HeaderComponent />
        <div class="contents">
        <div class="statistics">
            <h3 id="AllTime">All Time</h3><br><br>
            <button id="buttonDownloadLoto6Statistics" @click="downloadLoto6Statistics()">Download CSV</button><br><br>
            <p>Time : {{results[0].Time}}</p><br><br>
            <div class="statisticsTable">
                <table>
                    <thead>
                        <tr>
                            <th id="rank">Rank</th>
                            <th id="number">Number</th>
                            <th id="count">Count</th>
                            <th id="rate">Rate</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="row in results" :key="row.Number">
                            <td>{{row.Rank}}</td>
                            <td>{{row.Number}}</td>
                            <td>{{row.Count}}</td>
                            <td>{{Math.floor(row.Rate * 1000) / 1000}}%</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        <div class="latelyStatistics">
        <h3>Latest 10 Times</h3><br><br>
            <div class="latelyStatisticsTable">
                <table>
                    <thead>
                        <tr>
                            <th id="latelyRank">Rank</th>
                            <th id="latelyNumber">Number</th>
                            <th id="latelyCount">Count</th>
                            <th id="Time1">{{latelyResults[43].LatestTime1}}</th>
                            <th id="Time2">{{latelyResults[43].LatestTime2}}</th>
                            <th id="Time3">{{latelyResults[43].LatestTime3}}</th>
                            <th id="Time4">{{latelyResults[43].LatestTime4}}</th>
                            <th id="Time5">{{latelyResults[43].LatestTime5}}</th>
                            <th id="Time6">{{latelyResults[43].LatestTime6}}</th>
                            <th id="Time7">{{latelyResults[43].LatestTime7}}</th>
                            <th id="Time8">{{latelyResults[43].LatestTime8}}</th>
                            <th id="Time9">{{latelyResults[43].LatestTime9}}</th>
                            <th id="Time10">{{latelyResults[43].LatestTime10}}</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="row in limitDispStatistics" :key="row.Number">
                            <td>{{row.Rank}}</td>
                            <td>{{row.Number}}</td>
                            <td>{{row.Count}}</td>
                            <td>{{row.Time1}}</td>
                            <td>{{row.Time2}}</td>
                            <td>{{row.Time3}}</td>
                            <td>{{row.Time4}}</td>
                            <td>{{row.Time5}}</td>
                            <td>{{row.Time6}}</td>
                            <td>{{row.Time7}}</td>
                            <td>{{row.Time8}}</td>
                            <td>{{row.Time9}}</td>
                            <td>{{row.Time10}}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        </div>
    </div>
</template>

<script>
import HeaderComponent from "../modules/header.vue";
import axios from 'axios';
export default {
    name: 'Loto6StatisticsPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            results: "",
            latelyResults: "",
        }
    },
    mounted() {
        this.getLoto6Statistics()
        this.getLoto6LatelyStatistics()
    },
    computed: {
        limitDispStatistics() {
            this.convertResults()
            return this.latelyResults.slice(0, 43)
        }
    },
    methods: {
        async getLoto6Statistics() {
            const results = await axios.get("/getLoto6Statistics");
            this.results = results.data
        },
        async downloadLoto6Statistics() {
            axios.get("/downloadLoto6Statistics", { responseType: 'blob'})
                .then((res) => {
                    const url = window.URL.createObjectURL(new Blob([res.data]))
                    const link = document.createElement('a')
                    link.href = url
                    link.setAttribute('download', 'Loto6Statistics.csv')
                    document.body.appendChild(link)
                    link.click()
                    window.URL.revokeObjectURL(url)
                })
        },
        async getLoto6LatelyStatistics() {
            const latelyResults = await axios.get("/getLoto6LatelyStatistics");
            this.latelyResults = latelyResults.data
        },
        convertResults() {
            for (let i = 0; i < 43; i++) {
                if (this.latelyResults[i].Time1 == 1) {
                    this.latelyResults[i].Time1 = '○';
                } else {
                    this.latelyResults[i].Time1 = '×';
                }
                if (this.latelyResults[i].Time2 == 1) {
                    this.latelyResults[i].Time2 = '○';
                } else {
                    this.latelyResults[i].Time2 = '×';
                }
                if (this.latelyResults[i].Time3 == 1) {
                    this.latelyResults[i].Time3 = '○';
                } else {
                    this.latelyResults[i].Time3 = '×';
                }
                if (this.latelyResults[i].Time4 == 1) {
                    this.latelyResults[i].Time4 = '○';
                } else {
                    this.latelyResults[i].Time4 = '×';
                }
                if (this.latelyResults[i].Time5 == 1) {
                    this.latelyResults[i].Time5 = '○';
                } else {
                    this.latelyResults[i].Time5 = '×';
                }
                if (this.latelyResults[i].Time6 == 1) {
                    this.latelyResults[i].Time6 = '○';
                } else {
                    this.latelyResults[i].Time6 = '×';
                }
                if (this.latelyResults[i].Time7 == 1) {
                    this.latelyResults[i].Time7 = '○';
                } else {
                    this.latelyResults[i].Time7 = '×';
                }
                if (this.latelyResults[i].Time8 == 1) {
                    this.latelyResults[i].Time8 = '○';
                } else {
                    this.latelyResults[i].Time8 = '×';
                }
                if (this.latelyResults[i].Time9 == 1) {
                    this.latelyResults[i].Time9 = '○';
                } else {
                    this.latelyResults[i].Time9 = '×';
                }
                if (this.latelyResults[i].Time10 == 1) {
                    this.latelyResults[i].Time10 = '○';
                } else {
                    this.latelyResults[i].Time10 = '×';
                }
            }
        }
    }
};
</script>

<style>
.contents {
    display: flex;
}
.statisticsTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.statisticsTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
.statistics {
    margin: 50px;
}
.latelyStatistics {
    margin: 50px;
}
.latelyStatisticsTable th {
    border: 1px solid #000066;
    background: #66CCFF;
}
.latelyStatisticsTable td {
    border: 1px solid #000066;
    background: #ffffff;
}
#buttonDownloadLoto6Statistics {
    float: left;
}
p {
    float: left;
}
</style>
