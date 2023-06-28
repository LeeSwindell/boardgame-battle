import { useState, useNavigate, useEffect } from 'react';
import axios from 'axios';
import { logger } from '../logger/logger';

// const baseUrl = 'http://localhost:8000';
let lobbyUrl = import.meta.env.VITE_PROD_LOBBY_API_ENDPOINT;
const prodMode = import.meta.env.VITE_PROD_MODE;
if (prodMode === 'dev') {
  lobbyUrl = import.meta.env.VITE_DEV_LOBBY_API_ENDPOINT;
}
if (prodMode === 'run') {
  lobbyUrl = import.meta.env.VITE_LOBBY_URL;
}

function LoginPage({ onLogin }) {
  const [username, setUsername] = useState('');

  function handleSubmit(event) {
    event.preventDefault();
    axios
      .post(`${lobbyUrl}/login`, {
        username,
      })
      .then(() => {
        logger.log(`setting session id to ${username}`);
        localStorage.setItem('sessionid', username);
        onLogin(username);
      })
      .catch((error) => {
        logger.error('error sending /login post');
        logger.error(error);
      });
  }

  return (
    <div className="flex flex-col w-screen h-screen items-center justify-center">
      <form onSubmit={handleSubmit} className="flex flex-col items-center space-y-1">
        <label className="text-gray-700 font-bold mb-2" htmlFor="usernameInput">
          Username:
        </label>
        <input type="text" className="shadow border appearance-none rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline text-center" value={username} onChange={(e) => setUsername(e.target.value)} id="usernameInput" />
        <div className="items-center justify-center">
          <button type="submit" className="bg-blue-500 my-6 hover:bg-blue-700 shadow-sm text-white font-bold rounded p-4">Enter</button>
        </div>
      </form>
    </div>
  );
}

export default LoginPage;
