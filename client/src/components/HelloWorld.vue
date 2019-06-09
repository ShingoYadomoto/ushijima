<template>
  <button @click="showText">ステータスを表示する。</button>
  <div id="text">{{ text.status }}</div>
</template>

<script>
import axios from 'axios'
import vue2Dropzone from 'vue2-dropzone'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'


export default {
    data: function() {
        return {
            text: 'hello',
        }
    },
    methods: {
        showText: function() {
            axios.get('/').then(res => {
                this.text = res.data.items || {};
            }).catch(function (error) {
                console.log(error);
            });
        },
    }
}


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
