import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function CharSelect() {
    return (
        <select name="char" id="char-select" className="w-full p-2 rounded focus:ring-blue-500 focus:border-blue-500">
            <option value="">Select</option>
            <option value="Ron">Ron</option>
            <option value="Hermione">Hermione</option>
            <option value="Neville">Neville</option>
            <option value="Harry">Harry</option>
            <option value="Luna">Luna</option>
        </select>
    )
}

function nameInput() {
    return (
        <form action="">
            <input>name</input>
        </form>
    )
}

function Lobby() {
    let params = useParams()

    const [players, setPlayers] = useState([])
    const [socket, setSocket] = useState(null)

    useEffect(() => {
        const lobbyId = params["id"]

        if (socket === null && lobbyId) {
            // console.log(`creating new socket: ${socket}, ${lobbyId}`)
            const newSocket = new WebSocket(`ws://localhost:8000/lobby/${lobbyId}`)
        
            newSocket.onmessage = (event) => {
              setPlayers(JSON.parse(event.data).players)
            }
        
            setSocket(newSocket)

            return () => {
                if (newSocket.readyState === 1) {
                    newSocket.close()
                }
            }
        }
      }, [params])


    return (
        <div className="flex flex-col justify-center items-center">
            <h3 className="text-4xl text-transparent bg-clip-text font-extrabold bg-gradient-to-r from-blue-600 to-red-600 m-4">Wooooooooooooooooooo</h3>
            <div className="relative w-auto overflow-x-auto border shadow-md sm:rounded-lg m-4">
                <table className="table-auto text-left shadow-sm">
                    <thead className="bg-blue-500 text-white">
                    <tr >
                        <th className="px-6 py-2 font-bold">Name</th>
                        <th className="px-6 py-2 font-bold">Character</th>
                    </tr>
                    </thead>
                    <tbody className="">
                    {players.map((player, i) => {
                        if (i % 2 === 0) {
                        return (
                            <tr key={player.id} className="bg-gray-50 justify-center items-center px-4 py-2 rounded">
                            <td className="px-6 py-2">{player.name}</td>
                            <td className="px-6 py-2">
                                <CharSelect />
                            </td>
                            </tr>
                        )
                        } else {
                        return (
                            <tr key={player.id} className="justify-center items-center px-4 py-2 rounded">
                            <td className="px-6 py-2">{player.name}</td>
                            <td className="px-6 py-2">
                                <CharSelect />
                            </td>
                            </tr>
                        )
                        }
                    })}
                    </tbody>
                </table>
            </div>
        </div>
    )
}

export default Lobby;