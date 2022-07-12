<template>
  <div>
    <input type="text" placeholder="User name" v-model="username" />
    <input type="number" placeholder="Access type" v-model="accessType" />
    <button @click="createUser">Create</button>
  </div>
</template>
<script>
import { postRequest } from "@/axios/requests/postRequest";

export default {
  data() {
    return {
      username: "",
      accessType: 0,
    };
  },
  methods: {
    createUser() {
      let data = { username: this.username, accessType: this.accessType };
      postRequest(data, "/client/create").then((response) => {
        console.log(response.data);
      });
    },
  },
};
</script>
<style lang="scss"></style>
