<template>
  <div>
    <button @click="send">Hello</button>
  </div>
</template>
<script>
import { connect, sendMsg } from "@/websocket";
export default {
  mounted() {
    connect();
  },
  methods: {
    send() {
      let data = { color: "red", amount: 5000 };
      sendMsg(JSON.stringify(data));
    },
  },
};
</script>
<style lang="scss"></style>
