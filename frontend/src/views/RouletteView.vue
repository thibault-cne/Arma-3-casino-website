<template>
  <div>
    <WheelComponent :isSpinning="isSpinning" :outcome="outcome" />
    <RouletteInputComponent :newMessage="this.newMessage" />
    <button @click="roll">Click</button>
  </div>
</template>
<script>
import { connect, socket } from "@/websocket";
import WheelComponent from "@/components/rouletteComponents/WheelComponent.vue";
import RouletteInputComponent from "@/components/rouletteComponents/RouletteInputComponent.vue";
export default {
  components: {
    WheelComponent,
    RouletteInputComponent,
  },
  created() {
    connect();
    socket.onmessage = (msg) => {
      console.log(msg);
    };
  },
  data() {
    return {
      newMessage: "",
      isSpinning: false,
      outcome: 9,
    };
  },
  methods: {
    roll() {
      this.isSpinning = !this.isSpinning;
      setTimeout(() => {
        this.isSpinning = !this.isSpinning;
      }, 6 * 1000);
    },
  },
};
</script>
<style lang="scss"></style>
