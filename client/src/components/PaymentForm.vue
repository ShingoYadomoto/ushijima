<template>
    <div id="payment_form" class="container card-panel striped">
        <div v-show="isSuccess" class="payment_form__message card-panel blue lighten-1">保存しました。</div>
        <div v-if="!isSuccess && error" v-show="error" class="payment_form__message--error card-panel  pink accent-3">{{ error }}</div>

        <form method="POST">
            <div class="mui-select">
                <select name="month_id" v-model="params.monthId">
                    <option value="" disabled selected></option>
                    <option v-for="(month, idx) in monthList" :key="idx" :value="month.id">{{ month.display }}</option>
                </select>
                <label><i class="material-icons">event</i>Month</label>
            </div>

            <div class="mui-select">
                <select name="payment_type_id" v-model="params.paymentTypeId">
                    <option value="" disabled selected></option>
                    <option v-for="(paymentType, idx) in paymentTypeList" :key="idx" :value="paymentType.id">{{ paymentType.display }}</option>
                </select>
                <label><i class="material-icons">view_carousel</i>Payment Type</label>
            </div>

            <div class="mui-select">
                <select name="payment_status_id" v-model="params.paymentStatusId">
                    <option value="" disabled selected></option>
                    <option v-for="(paymentStatus, idx) in paymentStatusList" :key="idx" :value="paymentStatus.id">{{ paymentStatus.display }}</option>
                </select>
                <label><i class="material-icons">grade</i>Payment Status</label>
            </div>

            <div class="mui-textfield mui-textfield--float-label">
                <input  name="amount" type="number"  v-model="params.amount">
                <label><i class="material-icons">attach_money</i>Amount</label>
            </div>

            <button class="btn waves-effect waves-light" v-on:click="submit">Submit
                <i class="material-icons right">send</i>
            </button>
        </form>
    </div>
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
                monthList: {},
                paymentTypeList: {},
                paymentStatusList: {},
                params: {
                    monthId: null,
                    paymentTypeId: null,
                    paymentStatusId: null,
                    amount: null,
                },
                isSuccess: false,
                error: '',
            }
        },
        methods: {
            fetchMonths: function() {
                customAxios.get('http://localhost:8888/month').then(res => {
                    this.monthList = res.data.monthList;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            fetchPaymentTypes: function() {
                customAxios.get('http://localhost:8888/payment_type').then(res => {
                    this.paymentTypeList = res.data.paymentTypeList;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            fetchPaymentStatuses: function() {
                customAxios.get('http://localhost:8888/payment_status').then(res => {
                    this.paymentStatusList = res.data.paymentStatusList;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            submit: function(e) {
                e.preventDefault();

                const data = new FormData();
                data.append('month_id', this.params.monthId);
                data.append('payment_type_id', this.params.paymentTypeId);
                data.append('payment_status_id', this.params.paymentStatusId);
                data.append('amount', this.params.amount);

                customAxios.post('http://localhost:8888/payment/create', data).then(res => {
                    this.isSuccess = res.data.isSuccess;
                    this.error = res.data.error;
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        mounted() {
            this.fetchMonths();
            this.fetchPaymentTypes();
            this.fetchPaymentStatuses();
        }
    }
</script>

<style>
    #payment_form.container {
        width: 90%;
        padding: 10px;
        margin: auto;
    }
    #payment_form.container form{
        width: 50%;
        margin: auto;
    }
    #payment_form td {
        padding: 15px 10px;
    }
    #payment_form .material-icons {
        font-size: 16px;
        margin-right: 2px;
        vertical-align: middle;
    }
    .payment_form__message {
        margin-bottom: 20px;
        color: #fff;
        font-weight: bold;
    }
    .payment_form__message--error {
        color: #fff;
        font-weight: bold;
    }
</style>
