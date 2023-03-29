import axios from "axios";

const baseUrl = 'http://localhost:8000'
const api = axios.create({
  baseURL: baseUrl, // Replace with your API URL
  withCredentials: true, // Send cookies with every request
});

api.interceptors.request.use(config => {
  const sessionid = localStorage.getItem('sessionid');
  if (sessionid) {
    config.headers['Authorization'] = sessionid
  }
  return config;
});

export default api