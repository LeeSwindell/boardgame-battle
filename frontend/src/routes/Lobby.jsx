import { useEffect, useRef, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { api } from '../api';
import { logger } from '../logger/logger';

let socketUrl = import.meta.env.VITE_PROD_SOCKET_API;
const prodMode = import.meta.env.VITE_PROD_MODE;
if (prodMode === 'dev') {
  socketUrl = import.meta.env.VITE_DEV_SOCKET_API;
}

function CharSelect({ lobbyid, characterSelection, canEdit }) {
  const [character, setCharacter] = useState(characterSelection);

  useEffect(() => {
    if (canEdit) {
      api
        .post(`/lobby/${lobbyid}/setchar`, { character })
        .then(() => {
          // logger.log('char update');
        })
        .catch((res) => {
          logger.error(res);
        });
    }
  }, [character]);

  if (!canEdit) {
    return characterSelection;
  }

  const handleOptionChange = (event) => {
    setCharacter(event.target.value);
  };

  return (
    <form>
      <select name="char" id="char-select" value={characterSelection} onChange={handleOptionChange} className="block w-full p-2 rounded bg-gray-50 border border-gray-200 focus:ring-blue-500 focus:border-blue-500 ">
        <option value="Ron">Ron</option>
        <option value="Hermione">Hermione</option>
        <option value="Neville">Neville</option>
        <option value="Harry">Harry</option>
        <option value="Luna">Luna</option>
      </select>
    </form>
  );
}

function Lobby() {
  const params = useParams();
  const socket = useRef(null);
  const navigate = useNavigate();
  const [players, setPlayers] = useState([]);
  const [url, setUrl] = useState(null);

  useEffect(() => {
    if (url) {
      navigate(url);
    }
  });

  // Create the socket connection
  useEffect(() => {
    // socket.current = new WebSocket('ws://localhost:8000/connectsocket');
    if (socketUrl === 'localhost:8000') {
      socket.current = new WebSocket('ws://localhost:8000/connectsocket');
    } else {
      socket.current = new WebSocket(`wss://${socketUrl}/connectsocket`);
    }

    socket.current.onopen = () => logger.log('lobby socket opened');
    socket.current.onclose = () => logger.log('lobby socket closed');

    // cleanup socket connection and send a request to backend when leaving page.
    return () => {
      api.get(`/lobby/${params.id}/leave`);
      socket.current.close();
    };
  }, []);

  useEffect(() => {
    const lobbyId = params.id;

    api
      .get(`/lobby/${lobbyId}/refresh`)
      .then((res) => {
        // logger.log(res.data);
        setPlayers(res.data.players);
      })
      .catch(() => {
        logger.error('error refreshing lobbies');
        navigate('/');
      });

    if (socket !== null && lobbyId) {
      socket.current.onmessage = (event) => {
        const data = JSON.parse(event.data);
        switch (data.type) {
          case 'RefreshRequest':
            api
              .get(`/lobby/${lobbyId}/refresh`)
              .then((res) => {
                // logger.log(res.data.players);
                setPlayers(res.data.players);
              });
            break;
          case 'StartGame':
            setUrl(`/game/${lobbyId}`);
            break;
          default:
            break;
        }
      };
    }
  }, [params]);

  function startGame() {
    api
      .get(`/lobby/${params.id}/startgame`);
  }

  return (
    <div className="flex flex-col justify-center items-center">
      <h3 className="text-4xl text-transparent bg-clip-text font-extrabold bg-gradient-to-r from-blue-600 to-red-600 m-4">Wooooooooooooooooooo</h3>
      <div className="relative w-auto overflow-x-auto border shadow-md sm:rounded-lg m-4">
        <table className="table-auto text-left shadow-sm">
          <thead className="bg-blue-500 text-white">
            <tr>
              <th className="px-6 py-2 font-bold">Name</th>
              <th className="px-6 py-2 font-bold">Character</th>
            </tr>
          </thead>
          <tbody className="">
            {players.map((player) => (
              <tr key={player.id} className="justify-center items-center px-4 py-2 rounded border-b">
                <td className="px-4 py-2">{player.name}</td>
                <td className="px-4 py-2">
                  <CharSelect lobbyid={params.id} characterSelection={player.character} canEdit={player.name === localStorage.getItem('sessionid')} player={player.name} />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <button type="button" onClick={startGame} className="bg-blue-500 rounded py-2 px-4 m-4 text-white hover:bg-blue-700 hover:shadow-lg font-bold">Start</button>
    </div>
  );
}

export default Lobby;
