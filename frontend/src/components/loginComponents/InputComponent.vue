<template>
  <div>
    <input type="text" placeholder="username" v-model="username" />
    <input type="password" placeholder="password" v-model="password" />
    <button @click="submit" type="submit">Submit</button>
  </div>
</template>
<script>
import { postRequest } from "@/axios/requests/postRequest";
import { authStore } from "@/store/authStore";
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    submit() {
      let data = { username: this.username, password: this.password };

      postRequest(data, "/oauth/login", "json").then((response) => {
        let credentials = {
          accessToken: response.data.access_token,
          refreshToken: response.data.refresh_token,
        };

        authStore.commit("updateStorage", credentials);
      });
    },
  },
};
</script>
<style lang="css"></style>
