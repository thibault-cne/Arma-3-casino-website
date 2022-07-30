<template>
  <div>
    <div class="banner">
      <div class="input">
        <input class="amount-input" type="number" v-model="amount" />
        <button @click="clear">clear</button>
      </div>
      <div class="buttons">
        <button @click="addAmount(5)">+5</button>
        <button @click="addAmount(25)">+25</button>
        <button @click="addAmount(100)">+100</button>
        <button @click="multiplyAmount(0.5)">1/2</button>
        <button @click="multiplyAmount(2)">x2</button>
        <button>MAX</button>
      </div>
    </div>
    <div class="colorRow">
      <RoulettePlayersComponent
        :color="'red'"
        :playerAmount="this.amount"
        :players="this.redPlayers"
        :displayedText="'1 - 7'"
      />
      <RoulettePlayersComponent
        :color="'green'"
        :playerAmount="this.amount"
        :players="this.greenPlayers"
        :displayedText="'0'"
      />
      <RoulettePlayersComponent
        :color="'black'"
        :playerAmount="this.amount"
        :players="this.blackPlayers"
        :displayedText="'8 - 14'"
      />
    </div>
  </div>
</template>
<script>
import RoulettePlayersComponent from "./RoulettePlayersComponent.vue";
export default {
  components: {
    RoulettePlayersComponent,
  },
  props: {
    newMessage: String,
  },
  data() {
    return {
      amount: 0,
      redPlayers: {
        maxAmount: 0,
        players: [],
      },
      greenPlayers: {
        maxAmount: 0,
        players: [],
      },
      blackPlayers: {
        maxAmount: 0,
        players: [],
      },
    };
  },
  methods: {
    clear() {
      this.amount = 0;
    },
    addAmount(value) {
      this.amount += value;
    },
    multiplyAmount(value) {
      this.amount *= value;
    },
  },
};
</script>
<style scoped lang="scss">
.banner {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
  width: 100vw;
  height: 12vh;
  background-color: aqua;
}

.input {
  background-color: rebeccapurple;
  height: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.buttons {
  background-color: red;
  height: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.colorRow {
  margin-top: 4vh;
  display: flex;
  flex-direction: row;
  justify-content: space-around;
}
</style>
