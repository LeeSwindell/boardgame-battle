import Game from "./routes/Game"
import Home from "./routes/Home"
import Lobby from "./routes/lobby"
import Lobbies from "./routes/Lobbies"
import { Route, Routes } from "react-router-dom"

function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="lobby/:id" element={<Lobby />} />
      <Route path="lobbies" element={<Lobbies />} />
      <Route path="game" element={<Game />} />
    </Routes>
  )
}

export default App;