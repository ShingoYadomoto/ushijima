<template>
    <div>
        <button @click="showText">ステータスを表示する。</button>
        <div id="text">{{ text.status }}</div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    data: function() {
        return {
            text: 'hello',
        }
    },
    methods: {
        showText: function() {
            axios.get('http://localhost:8888/').then(res => {
                this.text = res.data.items[0].status || {};
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
