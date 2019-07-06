<template>
    <div id="app" class="ushijima_back">
        <mainContent></mainContent>
    </div>
</template>

<script>
import mainContent from './components/MainContent.vue'

export default {
  name: 'app',
  components: {
      mainContent
  }
}
</script>

<style>
#app {
  text-align: center;
  color: #2c3e50;
  padding-top: 230px;
}

.ushijima_back {
    display: block;
    min-height: 912px;
    background: url(./assets/images/ushijima_bottomback.jpg) no-repeat;
}
</style>