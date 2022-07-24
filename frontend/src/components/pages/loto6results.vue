<template>
  <div>
    <HeaderComponent />
    <button @click="downloadLoto6Results()">Download CSV</button><br><br>
    <div class="resultTable">
        <table>
            <thead>
                <tr>
                    <th id="time">Time</th>
                    <th colspan="6">Numbers</th>
                    <th>Edit</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="row in results" :key="row.Time">
                    <td>{{row.Time}}</td>
                    <td>{{row.Number_1}}</td>
                    <td>{{row.Number_2}}</td>
                    <td>{{row.Number_3}}</td>
                    <td>{{row.Number_4}}</td>
                    <td>{{row.Number_5}}</td>
                    <td>{{row.Number_6}}</td>
                    <td><button @click="moveEditLoto6Result(row.Time)">Edit</button></td>
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
    name: 'Loto6ResultPage',
    components: {
        HeaderComponent
    },
    data() {
        return {
            results:""
        }
    },
    mounted() {
        this.getLoto6Results()
    },
    methods: {
        async getLoto6Results() {
            const results = await axios.get("/getLoto6Results");
            this.results = results.data
        },
        async downloadLoto6Results() {
            axios.get("/downloadLoto6Results", { responseType: 'blob',})
                .then((res) => {
                    // Blobを参照するための一時的なURLを生成
                    const url = window.URL.createObjectURL(new Blob([res.data]))
                    // HTML要素のaタグを生成
                    const link = document.createElement('a')
                    link.href = url
                    // aタグのdownload属性を設定
                    link.setAttribute('download', 'Loto6Results.csv')
                    // 生成したaタグを設置し、クリックさせる
                    document.body.appendChild(link)
                    link.click()
                    // URLを削除
                    window.URL.revokeObjectURL(url)
                })
        },
        moveEditLoto6Result(time) {
            this.$router.push({
                name: 'editLoto6Result',
                params: { id: time}
            })
        } 
    }
};
</script>

<style>
.resultTable th,td {
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
</style>