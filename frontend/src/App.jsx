import { Route, Routes } from 'react-router-dom';
import { useState, useEffect } from 'react';
import Game from './routes/Game';
import Home from './routes/Home';
import Lobby from './routes/Lobby';
import Lobbies from './routes/Lobbies';
import LoginPage from './routes/LoginPage';
import api from './api';

function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [socket, setSocket] = useState(null);

  // create socket on component mount
  useEffect(() => {
    const sock = new WebSocket('ws://localhost:8000/connectsocket');
    setSocket(sock);

    return () => {
      if (sock.readyState === 1) {
        sock.close();
      }
    };
  }, []);

  // Check for existing user session
  useEffect(() => {
    const localSessionid = localStorage.getItem('sessionid');
    console.log('session id found locally:', localSessionid);

    api
      .get('/sessionid')
      .then((response) => {
        setIsLoading(false);
        if (response.data === true) {
          console.log('logged in with session id:', localSessionid);
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
      <Route path="/" element={<Home socket={socket} />} />
      <Route path="lobby/:id" loader={({ params }) => console.log(params[':id'])} element={<Lobby socket={socket} />} />
      <Route path="lobbies" element={<Lobbies socket={socket} />} />
      <Route path="game" element={<Game socket={socket} />} />
    </Routes>
  );
}

export default App;
