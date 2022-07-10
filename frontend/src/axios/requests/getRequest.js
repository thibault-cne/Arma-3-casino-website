import { getAPI } from "@/axios/axios";
import { refreshToken } from "@/axios/requests/refreshRequests";
import { createHeader } from "@/axios/requests/createHeader";

function getRequest(url, headerType, params = {}) {
  let header = createHeader(headerType);
  return new Promise((resolve, reject) => {
    getAPI
      .get(url, { headers: header, params: params })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 401) {
            refreshToken();
          }
        }
        reject(error);
      });
  });
}

export { getRequest };
