import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const baseUrl = 'http://localhost:8000'


function CharSelect() {
    return (
        <select name="char" id="char-select" className="w-full p-2 rounded bg-gray-50 focus:ring-blue-500 focus:border-blue-500">
            <option value="">Select</option>
            <option value="Ron">Ron</option>
            <option value="Hermione">Hermione</option>
            <option value="Neville">Neville</option>
            <option value="Harry">Harry</option>
            <option value="Luna">Luna</option>
        </select>
    )
}

function Lobby({ socket }) {
    let params = useParams()

    const [players, setPlayers] = useState([])

    // is the lobbyid param needed?
    useEffect(() => {
        const lobbyId = params["id"]

        axios.get(baseUrl+`/lobby/${lobbyId}/refresh`)
            .then((res) => {
                setPlayers(res.data.players)
            })

        if (socket !== null && lobbyId) {        
            socket.onmessage = (event) => {
                const data = JSON.parse(event.data)
                switch (data.type) {
                    case 'RefreshRequest':
                        axios.get(baseUrl+`/lobby/${lobbyId}/refresh`)
                            .then((res) => {
                                console.log(res.data)
                                setPlayers(res.data.players)
                            })
                    case 'LobbyUpdate':

                }
            }
        }
      }, [params])

    function handleRefresh() {
        const message = {
            'Type': 'RefreshLobby'
        }
        socket.send(JSON.stringify(message))
    }

    function addPlayer() {
        const newPlayer = {
            id: 123,
            name: 'Bing Bong',
            character: 'Ron',
          };

        axios
            .post(baseUrl+`/lobby/${params["id"]}/addplayer`, newPlayer)
            .then((res) => {
                console.log(res)
            })
            .catch((err) => {
                console.error(err)
            })
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
                    {players.map(player => {
                        return (
                            <tr key={player.id} className="justify-center items-center px-4 py-2 rounded border-b">
                            <td className="px-6 py-2">{player.name}</td>
                            <td className="px-6 py-2">
                                <CharSelect />
                            </td>
                            </tr>
                        )
                    })}
                    </tbody>
                </table>
            </div>
            <button className="border" onClick={handleRefresh}>refresh lobbies</button>
            <button className="border" onClick={addPlayer}>addPlayer</button>
        </div>
    )
}

export default Lobby;