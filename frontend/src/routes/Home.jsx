import { Link } from 'react-router-dom';
import axios from 'axios';
import api from '../api';

const baseUrl = 'http://localhost:8000';

function Home() {
  function createLobbyHandler() {
    console.log('post /lobby/create');

    api
      .get('/lobby/create')
      .then((response) => {

      });

    axios
      .get(`${baseUrl}/lobby/create`)
      .then((response) => {
        // do nothing?
      });
  }

  return (
    <div className="flex flex-col w-screen h-screen items-center p-8 space-y-24">
      <h3 className="text-4xl text-transparent bg-clip-text font-extrabold bg-gradient-to-r from-blue-600 to-red-600">Welcome to the Quantum Gaming Safari</h3>
      <div className="flex flex-row justify-between">
        <button to="/lobby/create" className="bg-blue-500 rounded p-4 m-4 text-white hover:bg-blue-700 hover:shadow-lg font-bold" onClick={createLobbyHandler}>Create Lobby</button>
        <Link to="/lobbies" className="bg-red-500 rounded p-4 m-4 text-white hover:bg-red-700 hover:shadow-lg font-bold">Join Lobby</Link>
      </div>
    </div>
  );
}

export default Home;
