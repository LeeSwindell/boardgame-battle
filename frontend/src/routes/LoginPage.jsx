import { useState } from 'react';
import axios from 'axios';

const baseUrl = 'http://localhost:8000';

function LoginPage({ onLogin }) {
  const [username, setUsername] = useState('');

  function handleSubmit(event) {
    event.preventDefault();
    axios
      .post(`${baseUrl}/login`, {
        username,
      })
      .then(() => {
        console.log(`setting session id to ${username}`);
        localStorage.setItem('sessionid', username);
        onLogin(username);
      })
      .catch((error) => {
        console.log('error sending /login post');
        console.log(error);
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
