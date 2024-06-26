import { Link, useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { api } from '../api';
import { logger } from '../logger/logger';

function Home() {
  const [url, setUrl] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (url) {
      navigate(url);
    }
  });

  function createLobbyHandler() {
    logger.log('get /lobby/create');

    api
      .get('/lobby/create')
      .then((response) => {
        const lobbyid = response.data;
        logger.log(lobbyid);
        const newUrl = `/lobby/${lobbyid}`;
        logger.log(newUrl);
        setUrl(newUrl);
      });
  }

  return (
    <div className="flex flex-col w-screen h-screen items-center p-8 space-y-24">
      <h3 className="text-4xl text-transparent bg-clip-text font-extrabold bg-gradient-to-r from-blue-600 to-red-600">Welcome to the Quantum Gaming Safari</h3>
      <div className="flex flex-row justify-between">
        <button type="submit" className="bg-blue-500 rounded p-4 m-4 text-white hover:bg-blue-700 hover:shadow-lg font-bold" onClick={createLobbyHandler}>Create Lobby</button>
        <Link to="/lobbies" className="bg-red-500 rounded p-4 m-4 text-white hover:bg-red-700 hover:shadow-lg font-bold">Join Lobby</Link>
      </div>
      {/* Put some text here to display to new users, updates and such? */}
      <div className="text-center py-16" />
    </div>
  );
}

export default Home;
