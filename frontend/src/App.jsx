import { Route, Routes, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import Game from './routes/Game';
import Home from './routes/Home';
import Lobby from './routes/Lobby';
import Lobbies from './routes/Lobbies';
import LoginPage from './routes/LoginPage';
import { api } from './api';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  // Check for existing user session
  useEffect(() => {
    // const localSessionid = localStorage.getItem('sessionid');
    // console.log('session id found locally:', localSessionid);

    api
      .get('/sessionid')
      .then((response) => {
        setIsLoading(false);
        if (response.data === true) {
          setLoggedIn(true);
        }
      })
      .catch((response) => {
        console.log('error using api to get /sessionid', response.data);
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
      <Route path="lobby/:id" loader={({ params }) => console.log(params[':id'])} element={<Lobby />} />
      <Route path="lobbies" element={<Lobbies />} />
      <Route path="game/:id" loader={({ params }) => console.log(params[':id'])} element={<Game />} />
    </Routes>
  );
}

export default App;
