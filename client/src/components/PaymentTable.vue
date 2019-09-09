<template>
    <table id="payment_table" class="container card-panel striped">
        <tbody>
            <tr>
                <th></th>
                <th v-for="(type, idx) in paymentTypeList" :key="idx">{{ type.display }}</th>
                <th>合計</th>
            </tr>
            <tr v-for="(ps, idx) in paymentsList" :key="idx">
                <td>{{ ps.month.display }}</td>
                <td v-for="p in ps.payments">
                    <i v-if="p.payment_status_id == 1" class="material-icons">check</i>
                    <i v-else-if="p.payment_status_id == 2" class="material-icons">check_box_outline_blank</i>
                    <i v-else-if="p.payment_status_id == 3" class="material-icons">sports_tennis</i>
                    <i v-else-if="p.payment_status_id == 4" class="material-icons">sports_soccer</i>

                    ¥{{ p.amount }}
                </td>
                <td>¥{{ ps.total_fee }}</td>
            </tr>
        </tbody>
    </table>
</template>

<script>
    import axios from 'axios'

    const customAxios = axios.create({
        auth: {
            username: 'user',
            password: 'password'
        }
    });

    export default {
        data: function() {
            return {
                paymentTypeList: {},
                paymentsList: {},
            }
        },
        methods: {
            fetchPaymentTypes: function() {
                customAxios.get('http://localhost:8888/payment_type').then(res => {
                    this.paymentTypeList = res.data.paymentTypeList;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            fetchMonths: function() {
                customAxios.get('http://localhost:8888/payment').then(res => {
                    this.paymentsList = res.data.paymentsList;
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        mounted() {
            this.fetchPaymentTypes();
            this.fetchMonths();
        }
    }
</script>

<style>
    #payment_table.container {
        width: 90%;
        padding: 10px;
        margin: auto;
    }
    #payment_table td {
        padding: 15px 10px;
    }
</style>
