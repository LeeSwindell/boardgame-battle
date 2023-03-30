import axios from 'axios';

const baseUrl = 'http://localhost:8000';
const api = axios.create({
  baseURL: baseUrl, // Replace with your API URL
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

export default api;
