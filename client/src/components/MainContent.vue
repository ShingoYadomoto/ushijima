<template>
    <div id="main_content">
        <PaymentTable></PaymentTable>
    </div>
</template>

<script>
import PaymentTable from "./PaymentTable";

export default {
    components: {PaymentTable},
}

</script>

<style>
    #main_content {
        background: white;
        height: 400px;
        display: flex;
    }
</style>
