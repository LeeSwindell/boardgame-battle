import { Route, Routes, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import Game from './routes/Game';
import Home from './routes/Home';
import Lobby from './routes/Lobby';
import Lobbies from './routes/Lobbies';
import LoginPage from './routes/LoginPage';
import { api } from './api';
import { logger } from './logger/logger';
import TestServer from './components/TestServer';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  // Check for existing user session
  useEffect(() => {
    api
      .get('/sessionid')
      .then((response) => {
        setIsLoading(false);
        if (response.data === true) {
          setLoggedIn(true);
        }
      })
      .catch((response) => {
        logger.error('error using api to get /sessionid', response.data);
      });
  }, []);

  if (isLoading) {
    return <div className="w-screen h-screen justify-center items-center">Loading...</div>;
  }

  if (!loggedIn) {
    return <LoginPage onLogin={() => setLoggedIn(true)} />;
  }

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="lobby/:id" loader={({ params }) => logger.log(params[':id'])} element={<Lobby />} />
      <Route path="lobbies" element={<Lobbies />} />
      <Route path="gamepage/:gameid" element={<Game />} />
      <Route path="testserver" element={<TestServer />} />
    </Routes>
  );
}

export default App;
