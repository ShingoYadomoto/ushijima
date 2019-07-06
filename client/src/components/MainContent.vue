<template>
    <div id="main_content">
        <a v-show="toggle" v-on:click="toggle = !toggle" class="waves-effect waves-light btn-large">Form<i class="material-icons right">chevron_right</i></a>
        <a v-show="!toggle" v-on:click="toggle = !toggle" class="waves-effect waves-light btn-large">Table<i class="material-icons right">chevron_right</i></a>
        <PaymentTable v-show="toggle"></PaymentTable>
        <PaymentForm v-show="!toggle"></PaymentForm>
    </div>
</template>

<script>
import PaymentTable from "./PaymentTable";
import PaymentForm from "./PaymentForm";

export default {
    components: {PaymentTable, PaymentForm},
    data: function() {
        return {
            toggle:true,
        }
    },}

</script>

<style>
    #main_content {
        background: white;
        height: 450px;
        display: flex;
    }
</style>
