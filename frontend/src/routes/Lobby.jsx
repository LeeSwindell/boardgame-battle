import { useEffect, useRef, useState } from 'react';
import { useParams } from 'react-router-dom';
import api from '../api';

function CharSelect({
  lobbyid, characterSelection, canEdit, player,
}) {
  const [character, setCharacter] = useState(characterSelection);

  useEffect(() => {
    console.log(`sending char post, char:${character} player:${player} charselection: ${characterSelection}`);

    if (canEdit) {
      api
        .post(`/lobby/${lobbyid}/setchar`, { character })
        .then(() => {
          console.log('char update');
        })
        .catch((res) => {
          console.log(res);
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
      <select name="char" id="char-select" value={characterSelection} onChange={handleOptionChange} className="block w-full p-2 rounded bg-gray-50 border border-gray-200 text-sm focus:ring-blue-500 focus:border-blue-500 ">
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
  const [players, setPlayers] = useState([]);
  const socket = useRef(null);

  // Create the socket connection
  useEffect(() => {
    socket.current = new WebSocket('ws://localhost:8000/connectsocket');
    socket.current.onopen = () => console.log('lobby socket opened');
    socket.current.onclose = () => console.log('lobby socket closed');
    return () => socket.current.close();
  }, []);

  useEffect(() => {
    const lobbyId = params.id;

    api
      .get(`/lobby/${lobbyId}/refresh`)
      .then((res) => {
        console.log(res.data);
        setPlayers(res.data.players);
      });

    if (socket !== null && lobbyId) {
      socket.current.onmessage = (event) => {
        const data = JSON.parse(event.data);
        switch (data.type) {
          case 'RefreshRequest':
            api
              .get(`/lobby/${lobbyId}/refresh`)
              .then((res) => {
                console.log(res.data.players);
                setPlayers(res.data.players);
              });
            break;
          default:
            break;
        }
      };
    }
  }, [params]);

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
    </div>
  );
}

export default Lobby;
