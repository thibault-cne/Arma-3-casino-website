import { refreshToken } from "@/axios/requests/refreshRequests";
import { authStore } from "@/store/authStore";

let uri = "ws://localhost:5454/api/v1/roulette/connect";

var socket = new WebSocket(uri + "?token=" + authStore.getters.accessToken);

let connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = (msg) => {
    let data = JSON.parse(msg.data);
    if (data.errorType === 401) {
      refreshToken().then(() => {
        connect();
      });
      return;
    }

    console.log(data);
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
