import axios from 'axios';

let lobbyUrl = import.meta.env.VITE_PROD_LOBBY_API_ENDPOINT;
let gameUrl = import.meta.env.VITE_PROD_GAME_API_ENDPOINT;
const prodMode = import.meta.env.VITE_PROD_MODE;
if (prodMode === 'dev') {
  lobbyUrl = import.meta.env.VITE_DEV_LOBBY_API_ENDPOINT;
  gameUrl = import.meta.env.VITE_DEV_GAME_API_ENDPOINT;
}
if (prodMode === 'run') {
  lobbyUrl = import.meta.env.VITE_LOBBY_URL;
  gameUrl = import.meta.env.VITE_GAME_URL;
}

const api = axios.create({
  baseURL: lobbyUrl,
  withCredentials: true, // Send cookies with every request
});

// Attach a user auth to each api request.
api.interceptors.request.use((config) => {
  const sessionid = localStorage.getItem('sessionid');
  if (sessionid) {
    // eslint-disable-next-line no-param-reassign
    config.headers.Authorization = sessionid;
  }
  return config;
});

const gameapi = axios.create({
  baseURL: gameUrl,
  withCredentials: true,
});

gameapi.interceptors.request.use((config) => {
  const sessionid = localStorage.getItem('sessionid');
  if (sessionid) {
    // eslint-disable-next-line no-param-reassign
    config.headers.Authorization = sessionid;
  }
  return config;
});

export { api, gameapi };
