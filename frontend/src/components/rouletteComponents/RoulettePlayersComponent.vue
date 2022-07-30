<template>
  <div class="color-row">
    <div class="color-btn" v-bind:class="this.color" @click="this.playerBet()">
      {{ this.displayedText }}
    </div>
  </div>
</template>
<script>
import { sendMsg } from "@/websocket";
export default {
  props: {
    playerAmount: Number,
    color: String,
    displayedText: String,
    players: {
      maxAmount: Number,
      players: [],
    },
  },
  methods: {
    playerBet() {
      if (this.playerAmount <= 0) {
        return;
      }

      let data = { color: this.color, amount: this.playerAmount };
      sendMsg(JSON.stringify(data));
    },
  },
};
</script>
<style scoped lang="scss">
.color-row {
  display: flex;
  align-items: center;
  flex-direction: column;
  width: 30%;
  background-color: aqua;
}
.color-btn {
  color: #ffffff;
  width: 100px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;

  &:hover {
    cursor: pointer;
  }
}

.red {
  background: #f95146;
}

.black {
  background: #2d3035;
}

.green {
  background: #00c74d;
}
</style>
