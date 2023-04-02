import { useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';
import api from '../api';

function Lobbies() {
  const [lobbies, setLobbies] = useState([]);
  const [url, setUrl] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (url) {
      navigate(url);
    }
  });

  useEffect(() => {
    api
      .get('/lobbies')
      .then((res) => {
        if (res.data) {
          setLobbies(res.data);
        }
      });
  }, []);

  function joinLobby(id) {
    return () => {
      api
        .get(`/lobby/${id}/join`)
        .then(() => {
          setUrl(`/lobby/${id}`);
        });
    };
  }

  return (
    <div className="flex flex-col w-screen h-screen items-center p-8 space-y-24">
      <h3 className="text-4xl text-transparent bg-clip-text font-extrabold bg-gradient-to-r from-blue-600 to-red-600">Open Lobbies</h3>
      <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
        <table className="table-auto text-left shadow-sm">
          <thead className="bg-blue-500 text-white">
            <tr>
              <th className="px-6 py-2 font-bold">Name</th>
              <th className="px-6 py-2 font-bold">Host</th>
              <th className="px-6 py-2 font-bold" aria-label="hidden" />
            </tr>
          </thead>
          <tbody className="">
            {lobbies.map((lobby, i) => {
              if (i % 2 === 0) {
                return (
                  <tr key={lobby.id} className="bg-gray-50 justify-center items-center px-4 py-2 rounded">
                    <td className="px-6 py-2">{lobby.name}</td>
                    <td className="px-6 py-2">{lobby.host}</td>
                    <td className="px-6 py-2">
                      <button onClick={joinLobby(lobby.id)} type="submit" className="bg-blue-500 hover:bg-blue-700 text-white justify-center items-center px-4 py-2 font-bold rounded">
                        Join
                      </button>
                    </td>
                  </tr>
                );
              }
              return (
                <tr key={lobby.id} className="justify-center items-center px-4 py-2 rounded">
                  <td className="px-6 py-2">{lobby.name}</td>
                  <td className="px-6 py-2">{lobby.host}</td>
                  <td className="px-6 py-2">
                    <button onClick={joinLobby(lobby.id)} type="submit" className="bg-blue-500 hover:bg-blue-700 text-white justify-center items-center px-4 py-2 font-bold rounded">
                      Join
                    </button>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Lobbies;
